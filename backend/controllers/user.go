package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	//Verificano se email é válido
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("Invalid email", err)
		c.Status(400)
		return
	}

	//Fazendo hash do password
	passwordUser, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		c.Status(400)
		log.Println("Error generate password", err)
	}

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

	secret := os.Getenv("SECRET")

	tokenString := c.GetHeader("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.Status(401)
		log.Println("Token JWT inválido ", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Status(400)
		log.Println("Erro ao obter claims do token JWT")
		return
	}

	sub, ok := claims["Issuer"].(float64)
	if !ok {
		c.Status(500)
		log.Println("Erro ao obter ID do usuario a partir do token JWT")
		return
	}
	userIDInt := int(sub)

	//Começando update
	bio := c.PostForm("bio")
	arroba := c.PostForm("arroba")

	foto, err := c.FormFile("foto_perfil")
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
			c.JSON(400, gin.H{
				"Error in upload": err.Error(),
			})
		}
	}

	arrobaData := ("@" + arroba)

	linkPerfil := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/perfil/%s", foto.Filename)
	linkCapa := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/perfil/%s", foto.Filename)

	query, err := DB.Query(`UPDATE usuario SET link_perfil = ?, link_capa = ?, bio = ?, arroba = ? WHERE id_usuario = ? `, linkPerfil, linkCapa, bio, arrobaData, userIDInt)
	if err != nil {
		log.Println("Error in update query", err)
		c.JSON(400, gin.H{
			"Erro in update query": err.Error(),
		})
		return
	}

	for query.Next() {
		err := query.Scan(&linkPerfil, &linkCapa, &bio, &arroba, &userIDInt)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.JSON(400, gin.H{
				"Error in scan result query": err.Error(),
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"link_foto_perfil": linkPerfil,
		"link_foto_capa":   linkCapa,
	})
}

func GetInfoUser(c *gin.Context) {
	secret := os.Getenv("SECRET")

	tokenString := c.GetHeader("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.Status(401)
		log.Println("Token JWT inválido ", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Status(400)
		log.Println("Erro ao obter claims do token JWT")
		return
	}

	sub, ok := claims["Issuer"].(float64)
	if !ok {
		c.Status(500)
		log.Println("Erro ao obter ID do usuario a partir do token JWT")
		return
	}
	userIDInt := int(sub)

	var getUser models.User

	query := DB.QueryRow(`SELECT nome, bio, arroba, link_perfil, link_capa FROM usuario WHERE id_usuario = ?`, userIDInt)

	if err := query.Scan(&getUser.Nome, &getUser.Bio, &getUser.Arroba, &getUser.LinkPerfil, &getUser.LinkCapa); err != nil {
		log.Println(err)
	}

	log.Println(getUser)
	c.JSON(200, getUser)
}
