package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Erro ao carregar variavel de ambiente", err)

	}

	secret := os.Getenv("SECRET")
	var credentials models.UserLogin
	var login = models.User{}

	err = c.ShouldBindJSON(&credentials)
	if err != nil {
		log.Println("Error in read body", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error in read body": err.Error(),
		})
		return
	}

	verificator := fmt.Sprintf(`SELECT id_usuario, nome, email, senha, dt_nascimento, bio FROM usuario WHERE email = '%s'`, credentials.Email)
	query, err := DB.Query(verificator)
	if err != nil {
		log.Println("Erro in select database", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro in select database": err.Error(),
		})
		return
	}

	for query.Next() {
		err := query.Scan(&login.IDUsuario, &login.Nome, &login.Email, &login.Senha, &login.DTNascimento, &login.Bio)
		if err != nil {
			log.Println("Erro in scan database", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro in scan database": err.Error(),
			})
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Senha), []byte(credentials.Senha))
	if err != nil {
		log.Println("Error in compare password", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro in compare password": err.Error(),
		})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer": login.IDUsuario,
		"Name":   login.Nome,
		"Bio":    login.Bio,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error generating jwt token", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error generating jwt token": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Logged": tokenString,
	})

}
