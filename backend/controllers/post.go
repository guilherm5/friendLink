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
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
)

func Post(c *gin.Context) {
	service := utils.UtilAWS()

	secret := os.Getenv("SECRET")
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Token JWT inválido",
		})
		log.Println("Token JWT inválido ", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao obter claims do token JWT",
		})
		log.Println("Erro ao obter claims do token JWT")
		return
	}

	sub, ok := claims["Issuer"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao obter ID do usuario a partir do token JWT",
		})
		log.Println("Erro ao obter ID do usuario a partir do token JWT")
		return
	}
	userIDInt := int(sub)

	legenda := c.PostForm("legenda_post")
	postTexto := c.PostForm("post_texto")

	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.Status(100)
		return
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			log.Println("Error opening file", err)
			c.Status(400)
			return
		}

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
			Key:    aws.String("post/" + file.Filename),
			Body:   src,
		}

		_, err = uploader.UploadWithContext(context.Background(), input)
		if err != nil {
			log.Println("Error in upload", err)
		}

	}
	linkPost := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/%s", file.Filename)

	query, err := DB.Query("INSERT INTO post (post_imagem, legenda_post, post_texto, id_usuario_pt) VALUES (?, ?, ?, ?)", linkPost, legenda, postTexto, userIDInt)
	if err != nil {
		log.Println("Error in post query", err)
		c.JSON(400, err)
		return
	}

	for query.Next() {
		err := query.Scan(&linkPost, &legenda, &postTexto, &userIDInt)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.JSON(400, err)
			return
		}
		c.JSON(200, "success")
	}
	c.JSON(200, gin.H{
		"link_image": linkPost,
	})
}

func Pagination(c *gin.Context) {
	var posts []models.Post
	var post models.Post
	var user models.User

	id, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro in convert integer", err)
		c.Status(400)
		return
	}
	limit := c.PostForm("limit_post")

	query := fmt.Sprintf(`SELECT p.id_post, p.legenda_post, p.post_imagem, p.post_texto, u.nome, u.arroba, u.link_perfil FROM post p INNER JOIN usuario u ON u.id_usuario = p.id_usuario_pt WHERE id_post > %v ORDER BY id_post LIMIT %v`, id, limit)

	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Error in pagination query", err)
		c.Status(400)
		return
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&post.IDPost, &post.LegendaPost, &post.PostImagem, &post.PostTexto, &user.Nome, &user.Arroba, &user.LinkPerfil)
		if err != nil {
			log.Println("Error in scanning row", err)
			c.Status(400)
			return
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in iterating rows", err)
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"pagination": posts,
	})
}
