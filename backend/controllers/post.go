package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
)

func Post(c *gin.Context) {
	service := utils.UtilAWS()

	strPost, err := uuid.NewV4()
	if err != nil {
		log.Println("Erro ao gerar uuid foto", err)
	}

	secret := os.Getenv("SECRET")
	tokenString := c.GetHeader("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.Status(400)
		log.Println("Token JWT inválido ", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Status(400)
		log.Println("Erro ao obter claims do token JWT", ok)
		return
	}

	sub, ok := claims["Issuer"].(float64)
	if !ok {
		c.Status(400)
		log.Println("Erro ao obter ID do usuario a partir do token JWT", ok)
		return
	}
	userIDInt := int(sub)

	postTexto := c.PostForm("post_texto")

	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Campo file null", err)
		c.Status(100)
		return
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			log.Println("Erro ao abrir arquivo", err)
			c.Status(400)
			return
		}

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

		_, err = DB.Exec("INSERT INTO post (post_imagem, post_texto, id_usuario_pt) VALUES (?, ?, ?)", linkPost, postTexto, userIDInt)
		if err != nil {
			log.Println("Erro ao executar query", err)
			c.JSON(400, err)
			return
		}
		c.JSON(200, gin.H{
			"link_image": linkPost,
		})
	}
}

func Pagination(c *gin.Context) {
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
	query := fmt.Sprintf(`SELECT p.id_post, p.post_imagem, p.post_texto, u.nome, u.arroba, u.link_perfil 
             FROM post p 
             INNER JOIN usuario u ON u.id_usuario = p.id_usuario_pt 
             %v
             ORDER BY id_post DESC 
             LIMIT %v`, where, limit)

	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Erro na query de paginação", err)
		c.Status(400)
		return
	}

	for rows.Next() {
		err := rows.Scan(&post.IDPost, &post.PostImagem, &post.PostTexto, &post.Nome, &post.Arroba, &post.LinkPerfil)
		if err != nil {
			log.Println("Erro ao scanear linhas", err)
			c.Status(400)
			return
		}
		posts = append(posts, post)
	}
	c.JSON(200, posts)
}
