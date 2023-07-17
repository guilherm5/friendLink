package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
)

func NewComment(c *gin.Context) {
	//var postComentario models.Comentario7
	IDUser := utils.GetUserJWT(c)

	comentario := c.PostForm("comentario")
	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO comentario (comentario, dt_comentario, id_post_cmt, id_usuario_cmt) VALUES (?, ?, ?, ?)`, comentario, time.Now(), IDPost, IDUser)
	if err != nil {
		log.Println("Erro ao comentar", err)
		c.Status(400)
	}

	c.Status(200)

}

func GetComment(c *gin.Context) {
	var comentario models.Comentario
	var getFullComent []models.Comentario

	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	query, err := DB.Query(`SELECT c.id_comentario, c.comentario, u.nome, u.link_perfil, p.id_post
	FROM comentario c 
	INNER JOIN usuario u ON u.id_usuario = c.id_usuario_cmt 
	INNER JOIN post p ON p.id_post = c.id_post_cmt
	WHERE p.id_post = ?`, IDPost)
	if err != nil {
		log.Println("Erro query Get comentario", err)
		c.Status(400)
		return
	}

	for query.Next() {
		if err := query.Scan(&comentario.IDComentario, &comentario.Comentario, &comentario.Nome, &comentario.LinkPerfil, &comentario.IDPost); err != nil {
			log.Println("Erro ao scanear comentarios do post", err)
			c.Status(400)
			return
		}
		getFullComent = append(getFullComent, comentario)

	}
	c.JSON(200, getFullComent)

}
