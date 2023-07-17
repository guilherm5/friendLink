package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/utils"
)

func NewCurtida(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)

	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO curtida (id_usuario_ct, id_post_ct, curtiu) VALUES (?, ?, ?)
	`, IDUser, IDPost, true)
	if err != nil {
		log.Println("Erro ao curtir", err)
		c.Status(400)
	}

	c.Status(200)

}

func DelCurtida(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)

	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`DELETE FROM curtida
	WHERE id_post_ct = ? AND id_usuario_ct = ?`, IDPost, IDUser)
	if err != nil {
		log.Println("Erro ao curtir", err)
		c.Status(400)
	}
	_, err = DB.Query(`UPDATE curtida
	SET curtiu = false
	WHERE id_usuario_ct = ?;
	`, IDUser)
	if err != nil {
		log.Println("Erro ao realizar update no campo 'curtiur'")
		c.Status(500)
	} else {
		c.Status(200)
	}

}
