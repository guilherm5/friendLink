package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Post(c *gin.Context) {
	err := godotenv.Load("./.env")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro ao acessar arquivo .env para carregar secret key": err,
		})
		log.Println("Erro ao acessar arquivo .env para carregar secret key ", err)
		return
	}

	secret := os.Getenv("SECRET")
	legenda := c.PostForm("legenda_post")

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
			"Erro ao obter ID do usuario a partir do token JWT": err,
		})
		log.Println("Erro ao obter ID do usuario a partir do token JWT", err)
		return
	}

	IDJwt := int(sub)

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error in read file", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in read file": err.Error(),
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		log.Println("Error open file", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error open file": err.Error(),
		})
		return
	}

	data, err := ioutil.ReadAll(src)
	if err != nil {
		if err != nil {
			log.Println("Error extract file file content ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error extract content ": err.Error(),
			})
			return
		}
	}

	query, err := DB.Query("INSERT INTO post (post_imagem, legenda_post, id_usuario_pt) VALUES (?, ?, ?)", data, legenda, IDJwt)
	if err != nil {
		log.Println("Error in query post Post", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in query post Post": err.Error(),
		})
		return
	}

	for query.Next() {
		err := query.Scan(&data, &legenda, &IDJwt)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Error in scan result query": err.Error(),
			})
		}
	}
}
