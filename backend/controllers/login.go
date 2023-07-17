package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	secret := os.Getenv("SECRET")
	//Variaveis para logar
	var credentials models.Usuario
	var login = models.Usuario{}

	//Aceitando creden. login via postform
	credentials.Email = c.PostForm("email")
	credentials.Senha = c.PostForm("senha")
	login.Senha = c.PostForm("senha")

	err := DB.QueryRow(`SELECT id_usuario, nome, senha FROM usuario WHERE email = ?`, credentials.Email).Scan(&login.IDUsuario, &login.Nome, &login.Senha)
	if err != nil {
		log.Println("Erro in scan database", err)
		c.Status(500)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Senha), []byte(credentials.Senha))
	if err != nil {
		log.Println("Error in compare password", err)
		c.Status(400)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer": login.IDUsuario,
		"Nome":   login.Nome,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error generating jwt token", err)
		c.Status(400)
	}

	c.JSON(http.StatusOK, gin.H{
		"logged": tokenString,
	})
}
