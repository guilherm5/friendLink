package utils

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GetUserJWT(c *gin.Context) (IDUser int) {
	//Capturando variavel de ambiente do servidor
	IDUser = 0
	if os.Getenv("env") != "production" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Println("Erro ao carregar variavel de ambiente para uso do middleware", err)
			c.Status(400)
			c.Abort()
		}
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Println("Secret middleware não pode ser null")
		c.Abort()
	}

	//setei meu secret em um context gin para não precisar recuperar com a mesma função em todo local..
	c.Set("secret", secret)

	authHeader := c.GetHeader("Authorization")
	token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		log.Println("Token inválido.", err)
		c.AbortWithStatus(401)
	}
	//capturando claims para setar id usuario em um context e resgatar em outros lugares
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Status(400)
		log.Println("Erro ao obter claims", err)
		c.Abort()
	}

	sub, ok := claims["Issuer"].(float64)
	if !ok {
		c.Status(500)
		log.Println("Erro ao obter ID do usuario a partir do token JWT", err)
		c.Abort()
	}
	IDUser = int(sub)
	c.Set("id", IDUser)
	return IDUser
}
