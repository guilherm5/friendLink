package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
)

func Post(c *gin.Context) {
	var lastID int
	var dataUser models.Pagination
	service := utils.UtilAWS()
	IDUser := utils.GetUserJWT(c)
	postTexto := c.PostForm("post_texto")

	strPost, err := uuid.NewV4()
	if err != nil {
		log.Println("Erro ao gerar uuid foto", err)
	}

	file, _ := c.FormFile("file")
	if err != nil {
		c.Status(200)
		c.Next()
	}

	if file == nil {
		_, err = DB.Exec("INSERT INTO post (post_texto, id_usuario_pt, dt_criacao) VALUES (?, ?, ?)", postTexto, IDUser, time.Now())
		if err != nil {
			log.Println("Erro ao executar query postexto", err)
			c.JSON(400, err)
			return
		}

		err = DB.QueryRow(`SELECT 
		p.id_post AS id_post,
		p.post_texto, 
		p.dt_criacao, 
		u.id_usuario,
		u.nome,
		u.arroba,
		u.link_perfil,
		COALESCE(COUNT(c.id_comentario), 0) AS qtde_comentario,
		COALESCE(COUNT(cur.id_post_ct), 0) AS qtde_curtida,
		(SELECT 
			CASE WHEN id_usuario_ct = 11 THEN 1 ELSE 0 END
		FROM curtida 
		WHERE id_post_ct = p.id_post AND id_usuario_ct = 11) AS curtiu
	FROM 
		post p 
	LEFT JOIN 
		comentario c ON c.id_post_cmt = p.id_post 
	LEFT JOIN 
		curtida cur ON cur.id_post_ct = p.id_post 
	INNER JOIN 
		usuario u ON u.id_usuario = p.id_usuario_pt
	WHERE 
		p.id_post = LAST_INSERT_ID()`).Scan(&lastID, &postTexto, &dataUser.DTCriacao, &dataUser.IDUser, &dataUser.Nome, &dataUser.Arroba, &dataUser.LinkPerfil, &dataUser.QtdeComentario, &dataUser.QtdeCurtida, &dataUser.Curtiu)
		if err != nil {
			log.Println("Erro ao executar query select last ID", err)
			c.JSON(400, err)
			return
		}

		c.JSON(200, gin.H{
			"post_imagem":     nil,
			"id_user":         IDUser,
			"post_texto":      postTexto,
			"id_post":         lastID,
			"nome":            dataUser.Nome,
			"arroba":          dataUser.Arroba,
			"link_perfil":     dataUser.LinkPerfil,
			"qtde_comentario": dataUser.QtdeComentario,
			"qtde_curtida":    dataUser.QtdeCurtida,
			"curtiu":          dataUser.Curtiu,
			"dt_criacao":      dataUser.DTCriacao,
		})
		c.Status(200)
	}
	src, _ := file.Open()

	_, err = io.ReadAll(src)
	if err != nil {
		log.Println("Erro ao ler arquivo", err)
		c.Status(400)
		return
	}
	src.Seek(0, io.SeekStart)

	uploader := s3manager.NewUploader(service)
	contentType := file.Header.Get("Content-Type")
	input := &s3manager.UploadInput{
		Bucket:             aws.String("frienlinkfotos"),
		Key:                aws.String("post/" + strPost.String()),
		Body:               src,
		ContentType:        &contentType,
		ContentDisposition: aws.String("inline"),
	}

	_, err = uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		log.Println("Erro ao realizar uplaod da imagem no bucket s3", err)
		c.Status(101)
	}
	linkPost := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/post/%s", strPost)

	_, err = DB.Exec("INSERT INTO post (post_imagem, post_texto, id_usuario_pt, dt_criacao) VALUES (?, ?, ?, ?)", linkPost, postTexto, IDUser, time.Now())
	if err != nil {
		log.Println("Erro ao executar query", err)
		c.JSON(400, err)
		return
	}

	err = DB.QueryRow(`SELECT 
		p.id_post AS id_post,
		p.post_texto,
		p.post_imagem, 
		p.dt_criacao,
		u.id_usuario, 
		u.nome,
		u.arroba,
		u.link_perfil,
		COALESCE(COUNT(c.id_comentario), 0) AS qtde_comentario,
		COALESCE(COUNT(cur.id_post_ct), 0) AS qtde_curtida,
		(SELECT 
			CASE WHEN id_usuario_ct = 11 THEN 1 ELSE 0 END
		FROM curtida 
		WHERE id_post_ct = p.id_post AND id_usuario_ct = 11) AS curtiu
	FROM 
		post p 
	LEFT JOIN 
		comentario c ON c.id_post_cmt = p.id_post 
	LEFT JOIN 
		curtida cur ON cur.id_post_ct = p.id_post 
	INNER JOIN 
		usuario u ON u.id_usuario = p.id_usuario_pt
	WHERE 
		p.id_post = LAST_INSERT_ID()`).Scan(&lastID, &postTexto, &linkPost, &dataUser.DTCriacao, &dataUser.IDUser, &dataUser.Nome, &dataUser.Arroba, &dataUser.LinkPerfil, &dataUser.QtdeComentario, &dataUser.QtdeCurtida, &dataUser.Curtiu)
	if err != nil {
		log.Println("Erro ao executar query select last ID", err)
		c.JSON(400, err)
		return
	}

	c.JSON(200, gin.H{
		"post_imagem":     linkPost,
		"id_user":         IDUser,
		"post_texto":      postTexto,
		"id_post":         lastID,
		"nome":            dataUser.Nome,
		"arroba":          dataUser.Arroba,
		"link_perfil":     dataUser.LinkPerfil,
		"qtde_comentario": dataUser.QtdeComentario,
		"qtde_curtida":    dataUser.QtdeCurtida,
		"curtiu":          dataUser.Curtiu,
	})
}

func Pagination(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)
	var posts []models.Pagination
	var post models.Pagination

	id, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}
	limit := c.PostForm("limit_post")

	var where = fmt.Sprintf("WHERE id_post < %d", id)
	if id == 0 {
		where = "WHERE 1=1"
	}

	query := fmt.Sprintf(`SELECT p.id_post, p.post_imagem, p.post_texto, p.dt_criacao, u.id_usuario, u.nome, u.arroba, u.link_perfil,
	COUNT(DISTINCT c.id_comentario) AS qtde_comentario,
	COUNT(DISTINCT cur.id_post_ct) AS qtde_curtida,
	(SELECT CASE 
		WHEN id_usuario_ct = %v THEN 1 
		ELSE 0 
		END
	FROM curtida 
		WHERE id_post_ct = p.id_post AND id_usuario_ct = %v) AS curtiu,

		(SELECT CASE 
			WHEN id_usuario_segue = %v THEN 1 
			ELSE 0 
			END
		FROM seguidor s  
			WHERE id_usuario_seguindo = u.id_usuario AND id_usuario_segue = %v) AS seguindo
	FROM post p
		INNER JOIN usuario u ON u.id_usuario = p.id_usuario_pt
		LEFT JOIN comentario c ON c.id_post_cmt =  p.id_post 
		LEFT JOIN curtida cur ON p.id_post = cur.id_post_ct 
	%s
	GROUP BY p.id_post
	ORDER BY p.id_post DESC
	LIMIT %s;`, IDUser, IDUser, IDUser, IDUser, where, limit)

	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Erro na query de paginação", err)
		c.Status(400)
		return
	}

	for rows.Next() {
		err := rows.Scan(&post.IDPost, &post.PostImagem, &post.PostTexto, &post.DTCriacao, &post.IDUser, &post.Nome, &post.Arroba, &post.LinkPerfil, &post.QtdeComentario, &post.QtdeCurtida, &post.Curtiu, &post.Seguindo)
		if err != nil {
			log.Println("Erro ao scanear linhas", err)
			c.Status(400)
			return
		}
		posts = append(posts, post)
	}
	c.JSON(200, posts)
}

func DeletePost(c *gin.Context) {
	//Funcao para pegar id do usuario
	IDUser := utils.GetUserJWT(c)

	IDComentario, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`DELETE FROM post WHERE id_post = ? AND id_usuario_pt = ?`, IDComentario, IDUser)
	if err != nil {
		log.Println("Erro ao deletar post", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}
