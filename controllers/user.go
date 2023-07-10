package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/database"
	"github.com/guilherm5/utils"
	"golang.org/x/crypto/bcrypt"
)

// Iniciando Banco de Dados
var DB = database.Init()

func NewUser(c *gin.Context) {
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	senha := c.PostForm("senha")

	//Verificano se email é válido
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("Invalid email", err)
		c.Status(400)
		return
	}

	//Fazendo hash do password
	passwordUser, err := bcrypt.GenerateFromPassword([]byte(senha), 14)

	//Preparando insert (resolvi usar para fins de teste, mas poso somente executra direto o insert)
	query, err := DB.Prepare(`INSERT INTO usuario (nome, email, senha) VALUES (?,?,?)`)
	if err != nil {
		log.Println("Error in prepare query insert", err)
		c.Status(500)
		return
	} else {
		//Executando insert
		query.Exec(nome, email, passwordUser)
		c.Status(201)
	}
}

func InfoUser(c *gin.Context) {
	//Configurando aws para receber imagem
	service := utils.UtilAWS()

	//Recuperando usuario pelo JWT
	IDJwt := c.GetInt("id")

	//Começando update
	bio := c.PostForm("bio")
	arroba := c.PostForm("arroba")

	foto, err := c.FormFile("foto")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.Status(100)
		return
	}

	if foto != nil {
		src, err := foto.Open()
		if err != nil {
			log.Println("Error opening file", err)
			c.Status(400)
			return
		}
		defer src.Close()

		_, err = io.ReadAll(src)
		if err != nil {
			log.Println("Error reading file content", err)
			c.Status(400)
			return
		}
		src.Seek(0, io.SeekStart)

		uploader := s3manager.NewUploader(service)

		input := &s3manager.UploadInput{
			Bucket: aws.String("frienlinkfotos"),
			Key:    aws.String("perfil/" + foto.Filename),
			Body:   src,
		}

		_, err = uploader.UploadWithContext(context.Background(), input)
		if err != nil {
			log.Println("Error in upload", err)
		}

	}

	fotoCapa, err := c.FormFile("foto_capa")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.Status(100)
		return
	}

	if fotoCapa != nil {
		srcCapa, err := fotoCapa.Open()
		if err != nil {
			log.Println("Error opening file", err)
			c.Status(400)
			return
		}

		_, err = io.ReadAll(srcCapa)
		if err != nil {
			log.Println("Error reading file content", err)
			c.Status(400)
			return
		}
		_, err = io.ReadAll(srcCapa)
		if err != nil {
			log.Println("Error reading file content", err)
			c.Status(400)
			return
		}
		srcCapa.Seek(0, io.SeekStart)

		inputCapa := &s3manager.UploadInput{
			Bucket: aws.String("frienlinkfotos"),
			Key:    aws.String("perfil/" + fotoCapa.Filename),
			Body:   srcCapa,
		}

		uploaderCapa := s3manager.NewUploader(service)
		_, err = uploaderCapa.UploadWithContext(context.Background(), inputCapa)
		if err != nil {
			log.Println("Error in upload", err)
		}
	}

	arrobaData := ("@" + arroba)

	linkPerfil := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/%s", foto.Filename)
	linkCapa := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/%s", fotoCapa.Filename)

	query, err := DB.Query(`UPDATE usuario SET link_perfil = ?, link_capa = ?, bio = ?, arroba = ? WHERE id_usuario = ? `, linkPerfil, linkCapa, bio, arrobaData, IDJwt)
	if err != nil {
		log.Println("Error in update query", err)
		c.Status(500)
		return
	}

	for query.Next() {
		err := query.Scan(&linkPerfil, &linkCapa, &bio, &arroba)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.Status(500)
			return
		}
	}
}
