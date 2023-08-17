package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/guilherm5/database"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
	"golang.org/x/crypto/bcrypt"
)

// Iniciando Banco de Dados
var DB = database.Init()

func NewUser(c *gin.Context) {
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	senha := c.PostForm("senha")
	if nome == "" || email == "" || senha == "" {
		log.Println("Campo nome, email ou senha não pode ser null")
		c.Status(400)
		return
	}

	//Verificano se email é válido
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("Email invalido", err)
		c.Status(400)
		return
	}

	//Fazendo hash do password
	passwordUser, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		c.Status(500)
		log.Println("Erro ao gerar hash da senha", err)
		return
	}

	//Executando insert
	_, err = DB.Exec(`INSERT INTO usuario (nome, email, senha, dt_criacao) VALUES (?,?,?,?)`, nome, email, passwordUser, time.Now())
	if err != nil {
		log.Println("Erro ao executar insert", err)
		c.Status(500)
		return
	} else {
		c.Status(201)
	}
}

func AfterCad(c *gin.Context) {
	//Configurando aws para receber imagem
	service := utils.UtilAWS()
	IDUser := utils.GetUserJWT(c) //Funcao para pegar id do usuario

	//Gerando uuid unico para cada foto/capa
	strFoto, err := uuid.NewV4()
	if err != nil {
		log.Println("Erro ao gerar uuid foto", err)
		c.Status(100)
	}

	strCapa, err := uuid.NewV4()
	if err != nil {
		log.Println("Erro ao gerar uuid capa", err)
		c.Status(100)
	}

	//Começando update
	bio := c.PostForm("bio")
	arroba := c.PostForm("arroba")

	foto, err := c.FormFile("foto_perfil")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Arquivo inexistente", err)
		c.Status(100)
		return
	}

	if foto != nil {
		src, err := foto.Open()
		if err != nil {
			log.Println("Erro ao abrir arquivo", err)
			c.Status(400)
			return
		}

		defer src.Close()

		_, err = io.ReadAll(src)
		if err != nil {
			log.Println("Erro ao ler conteudo do arquivo", err)
			c.Status(400)
			return
		}
		src.Seek(0, io.SeekStart)

		//upload file no bucket s3
		uploader := s3manager.NewUploader(service)
		contentType := foto.Header.Get("Content-Type")
		input := &s3manager.UploadInput{
			Bucket:             aws.String("frienlinkfotos"),
			Key:                aws.String("perfil/" + strFoto.String()),
			Body:               src,
			ContentType:        &contentType,
			ContentDisposition: aws.String("inline"),
		}

		_, err = uploader.UploadWithContext(context.Background(), input)
		if err != nil {
			log.Println("Erro ao realizar upload do arquivo no s3", err)
			c.Status(500)
		}

	}

	fotoCapa, err := c.FormFile("foto_capa")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Arquivo inexistente", err)
		c.Status(100)
		return
	}

	if fotoCapa != nil {
		srcCapa, err := fotoCapa.Open()
		if err != nil {
			log.Println("Erro ao abrir arquivo", err)
			c.Status(400)
			return
		}

		_, err = io.ReadAll(srcCapa)
		if err != nil {
			log.Println("Erro ao ler conteudo do arquivo", err)
			c.Status(400)
			return
		}
		srcCapa.Seek(0, io.SeekStart)

		//upload file no bucket s3
		contentTypeCapa := foto.Header.Get("Content-Type")
		inputCapa := &s3manager.UploadInput{
			Bucket:             aws.String("frienlinkfotos"),
			Key:                aws.String("perfil/" + strCapa.String()),
			Body:               srcCapa,
			ContentType:        &contentTypeCapa,
			ContentDisposition: aws.String("inline"),
		}

		uploaderCapa := s3manager.NewUploader(service)
		_, err = uploaderCapa.UploadWithContext(context.Background(), inputCapa)
		if err != nil {
			log.Println("Error ao realizar upload da foto para o s3", err)
			c.Status(500)
		}
	}

	//tratando strings de arroba e link de fotos
	arrobaData := ("@" + arroba)
	linkPerfil := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/perfil/%s", strFoto)
	linkCapa := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/perfil/%s", strCapa)

	_, err = DB.Exec(`UPDATE usuario SET link_perfil = ?, link_capa = ?, bio = ?, arroba = ? WHERE id_usuario = ? `, linkPerfil, linkCapa, bio, arrobaData, IDUser)
	if err != nil {
		log.Println("Error ao executar update", err)
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"link_foto_perfil": linkPerfil,
		"link_foto_capa":   linkCapa,
	})
}

func GetInfoUser(c *gin.Context) {
	//funcao que retorna id do usuario
	IDUser := utils.GetUserJWT(c)
	var getUser models.Usuario

	query := DB.QueryRow(`SELECT id_usuario, nome, bio, arroba, link_perfil, link_capa, dt_criacao FROM usuario WHERE id_usuario = ?`, IDUser)
	if err := query.Scan(&getUser.IDUsuario, &getUser.Nome, &getUser.Bio, &getUser.Arroba, &getUser.LinkPerfil, &getUser.LinkCapa, &getUser.DataCriacao); err != nil {
		log.Println(err)
	}

	c.JSON(200, getUser)
}

// pegar informações para o feed do usuario
func FeedUser(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)
	var getUser models.FeedUser
	var dataUser []models.FeedUser

	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Errop ao converter id post", err)
		c.Status(400)
	}

	limit := c.PostForm("limit")

	query, err := DB.Query(`SELECT u.id_usuario, u.nome, u.arroba, u.link_perfil, u.link_capa, u.dt_criacao,
	p.id_post, p.post_texto, p.post_imagem, 
	COUNT(DISTINCT c.id_comentario) AS qtde_comentario, 
	COUNT(DISTINCT cur.id_post_ct) AS qtde_curtida,
	(SELECT CASE 
			WHEN id_usuario_ct = ? THEN 1 
			ELSE 0 
			END
		FROM curtida 
			WHERE id_post_ct = p.id_post AND id_usuario_ct = ?) AS curtiu
	FROM usuario u
	INNER JOIN post p ON p.id_usuario_pt = u.id_usuario 
	LEFT JOIN comentario c ON c.id_post_cmt = p.id_post 
	LEFT JOIN curtida cur ON cur.id_post_ct = p.id_post 
	WHERE u.id_usuario = ?
	AND p.id_post < ?
	GROUP BY p.id_post
	ORDER BY p.id_post DESC
	LIMIT ?`, IDUser, IDUser, IDUser, IDPost, limit)
	if err != nil {
		log.Println("Erro ao buscar posts do usuario", err)
	}

	for query.Next() {
		err := query.Scan(&getUser.IDUsuario, &getUser.Nome, &getUser.Arroba, &getUser.LinkPerfil, &getUser.LinkCapa, &getUser.DataCriacao, &getUser.IDPost, &getUser.PostTexto, &getUser.PostImagem, &getUser.QtdeComentario, &getUser.QtdeCurtida, &getUser.Curtiu)
		if err != nil {
			log.Println("Erro ao scanear linhas", err)
			c.Status(400)
			return
		}
		dataUser = append(dataUser, getUser)
	}

	c.JSON(200, dataUser)
}
