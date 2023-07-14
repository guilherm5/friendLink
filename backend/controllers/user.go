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
	"github.com/gofrs/uuid"
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
		c.Status(400)
		log.Println("Erro ao gerar hash da senha", err)
		return
	}

	//Executando insert
	_, err = DB.Exec(`INSERT INTO usuario (nome, email, senha) VALUES (?,?,?)`, nome, email, passwordUser)
	if err != nil {
		log.Println("Erro ao executar insert", err)
		c.Status(400)
		return
	} else {
		c.Status(201)
	}

}

func AfterCad(c *gin.Context) {
	//Configurando aws para receber imagem
	service := utils.UtilAWS()

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

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Println("Secret não pode ser null")
		c.Status(500)
	}

	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		log.Println("Token JWT inválido ", err)
		c.Status(401)
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
		log.Println("Erro ao obter ID do usuario a partir do token JWT", ok)
		c.Status(500)
		return
	}
	userIDInt := int(sub)

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

	_, err = DB.Exec(`UPDATE usuario SET link_perfil = ?, link_capa = ?, bio = ?, arroba = ? WHERE id_usuario = ? `, linkPerfil, linkCapa, bio, arrobaData, userIDInt)
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
