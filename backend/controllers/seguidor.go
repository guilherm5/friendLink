package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/utils"
)

func Seguir(c *gin.Context) {
	IDUserSegue := utils.GetUserJWT(c)

	IDSeguindo, err := strconv.Atoi(c.PostForm("id_usuario_seguindo"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO seguidor (id_usuario_segue, id_usuario_seguindo) VALUES (?, ?)`, IDUserSegue, IDSeguindo)
	if err != nil {
		log.Println("Erro ao seguir usuario", err)
		c.Status(400)
	}
	c.Status(200)
}

func Desseguir(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)

	IDUsuarioSeguindo, err := strconv.Atoi(c.PostForm("id_usuario_seguindo"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`DELETE FROM seguidor
	WHERE id_usuario_segue = ? AND id_usuario_seguindo = ?`, IDUser, IDUsuarioSeguindo)
	if err != nil {
		log.Println("Erro ao deletar seguidor", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}
