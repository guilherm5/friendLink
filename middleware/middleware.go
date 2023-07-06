package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// Middleware de verificação (login)
func MiddlewareGO() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Capturando variavel de ambiente do servidor
		if os.Getenv("env") != "production" {
			err := godotenv.Load("./.env")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"erro ao carregar variavel de ambiente para middleware": err,
				})
				log.Println("erro ao carregar variavel de ambiente para middleware", err)
			}
		}
		secret := os.Getenv("SECRET")
		//setei meu secret em um context gin para não precisar recuperar com a mesma função em todo local..
		c.Set("secret", secret)

		authHeader := c.GetHeader("Authorization")
		token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Token invalido ou expirado": err,
			})
			log.Println("Token invalido ou expirado", err)
			return
		}
		//capturando claims para setar id usuario em um context e resgatar em outros lugares
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
		c.Set("id", IDJwt)
		c.Next()

	}
}
