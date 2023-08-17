package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/utils"
)

func NewCurtida(c *gin.Context) {
	//Funcao para pegar id do usuario
	IDUser := utils.GetUserJWT(c)

	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO curtida (id_usuario_ct, id_post_ct) VALUES (?, ?)
	`, IDUser, IDPost)
	if err != nil {
		log.Println("Erro ao curtir", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}

func DelCurtida(c *gin.Context) {
	//Funcao para pegar id do usuario
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
		log.Println("Erro ao deletar curtida do post", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}

func CurteComment(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)

	IDComment, err := strconv.Atoi(c.PostForm("id_comentario"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO curt_comentario(id_usuario_curt, id_comentario_curt) VALUES (?,?)`, IDUser, IDComment)
	if err != nil {
		log.Println("Erro ao curtir comentario", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}

func CurtRespComment(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)

	IDComment, err := strconv.Atoi(c.PostForm("id_resp_comentario"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO curt_resp_comentario(id_usuario_curt_resp, id_curt_resp_comentario) VALUES (?,?)`, IDUser, IDComment)
	if err != nil {
		log.Println("Erro ao curtir resposta comentario", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}

func DelCurtidaComentario(c *gin.Context) {
	//Funcao para pegar id do usuario
	IDUser := utils.GetUserJWT(c)

	IDComentario, err := strconv.Atoi(c.PostForm("id_comentario"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`DELETE FROM curt_comentario WHERE id_comentario_curt = ? AND id_usuario_curt = ?`, IDComentario, IDUser)
	if err != nil {
		log.Println("Erro ao deletar curtida do comentario", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}

func DelCurtidaRespComentario(c *gin.Context) {
	//Funcao para pegar id do usuario
	IDUser := utils.GetUserJWT(c)

	IDCOmentarioResp, err := strconv.Atoi(c.PostForm("id_curt_resp_comentario"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`DELETE FROM curt_resp_comentario WHERE id_curt_resp_comentario = ? AND id_usuario_curt_resp = ?`, IDCOmentarioResp, IDUser)
	if err != nil {
		log.Println("Erro ao deletar curtida do comentario", err)
		c.Status(400)
	} else {
		c.Status(200)
	}
}
