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
		c.Status(400)
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

	claimsRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer": login.IDUsuario,
		"Nome":   login.Nome,
		"exp":    time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenRefresh, err := claimsRefresh.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error generating jwt token", err)
		c.Status(400)
	}

	c.JSON(http.StatusOK, gin.H{
		"logged":  tokenString,
		"refresh": tokenRefresh,
	})
}

func Refresh(c *gin.Context) {
	secret := os.Getenv("SECRET")
	RefreshTokenString := c.GetHeader("Authorization")

	refreshToken, err := jwt.Parse(RefreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("Refresh token invalido", err)
		c.AbortWithStatus(401)
		return
	}
	if !refreshToken.Valid {
		c.AbortWithStatus(400)
		return
	}
	claims := refreshToken.Claims.(jwt.MapClaims)

	Nome := claims["Nome"].(string)
	Issuer := int64(claims["Issuer"].(float64))

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Nome":   Nome,
		"Issuer": Issuer,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	TokenString, err := Token.SignedString([]byte(secret))
	if err != nil {
		c.Status(500)
		log.Println("Erro ao gerar novo token JWT", err)
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"new_token": TokenString,
	})
}
