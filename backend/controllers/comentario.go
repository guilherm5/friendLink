package controllers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
)

func NewComment(c *gin.Context) {
	//funcao para pegar id do usuario
	IDUser := utils.GetUserJWT(c)
	var comment = models.Comentario{}

	comment.IDUsuarioCmt = IDUser
	comment.Comentario = c.PostForm("comentario")
	comment.IDPost, _ = strconv.Atoi(c.PostForm("id_post"))

	_, err := DB.Exec(`INSERT INTO comentario (comentario, dt_comentario, id_post_cmt, id_usuario_cmt) VALUES (?, ?, ?, ?); `, comment.Comentario, time.Now(), comment.IDPost, comment.IDUsuarioCmt)
	if err != nil {
		log.Println("Erro ao comentar", err)
		c.Status(400)
	}

	err = DB.QueryRow(`
	SELECT u.nome, u.arroba, u.link_perfil, c.dt_comentario, c.id_comentario 
	FROM comentario AS c
	INNER JOIN usuario u ON u.id_usuario = c.id_usuario_cmt
	WHERE c.id_comentario = LAST_INSERT_ID()
	`).Scan(&comment.Nome, &comment.Arroba, &comment.LinkPerfil, &comment.DTComentario, &comment.IDComentario)

	if err != nil {
		log.Println("Erro ao obter arroba do usuário", err)
		c.Status(500)
		return
	}

	c.JSON(200, comment)
}

func GetComment(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)
	var comentario models.Comentario
	var getFullComent []models.Comentario

	limit, err := strconv.Atoi(c.PostForm("limit"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}
	IDPost, err := strconv.Atoi(c.PostForm("id_post"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}
	IDPage, err := strconv.Atoi(c.PostForm("id_comentario"))
	if err != nil {
		log.Println("Erro ao converter id", err)
		c.Status(400)
		return
	}

	var and = fmt.Sprintf("AND c.id_comentario < %d", IDPage)
	if IDPage == 0 {
		and = "AND 1=1"
	}

	selector := fmt.Sprintf(`SELECT c.id_comentario, c.dt_comentario, c.comentario, c.id_usuario_cmt,
	u.nome, u.arroba, u.link_perfil, p.id_post,
	COUNT(DISTINCT rc.id_resp_comentario) AS qtde_resposta_comentario,
	COUNT(DISTINCT cc.id_usuario_curt) AS qtde_curtida_comentario,
	(SELECT CASE 
		WHEN id_usuario_curt = %d THEN 1
		ELSE 0
	END
	FROM curt_comentario 
		WHERE id_comentario_curt = c.id_comentario
		AND id_usuario_curt = %d
	) AS curtiu_comentario,
	COUNT(crc.id_curt_resp_comentario) AS qtde_curtida_resp_comentario,
	(SELECT CASE 
	WHEN id_usuario_curt_resp = %d THEN 1 
	ELSE 0
	END
	FROM curt_resp_comentario 
	WHERE id_curt_resp_comentario = rc.id_resp_comentario
	AND id_usuario_curt_resp = %d
	) AS curtiu_resp_comentario
	FROM comentario c 
	INNER JOIN usuario u				ON u.id_usuario					= c.id_usuario_cmt 
	INNER JOIN post p					ON p.id_post					= c.id_post_cmt
	LEFT JOIN resp_comentario rc 		ON rc.id_comentario_rp			= c.id_comentario
	LEFT JOIN curt_comentario cc		ON cc.id_comentario_curt		= c.id_comentario
	LEFT JOIN curt_resp_comentario crc	ON crc.id_curt_resp_comentario	= rc.id_resp_comentario
	WHERE p.id_post = %d
	%s
	GROUP BY c.id_comentario 
	ORDER BY c.id_comentario DESC
	LIMIT %d`, IDUser, IDUser, IDUser, IDUser, IDPost, and, limit)

	query, err := DB.Query(selector)
	if err != nil {
		log.Println("Erro query Get comentario", err)
		c.Status(400)
		return
	}

	for query.Next() {
		if err := query.Scan(&comentario.IDComentario, &comentario.DTComentario, &comentario.Comentario, &comentario.IDUsuarioCmt, &comentario.Nome, &comentario.Arroba, &comentario.LinkPerfil, &comentario.IDPost, &comentario.QtdeRespostaComentario, &comentario.QtdeCurtidaComentario, &comentario.CurtiuComentario, &comentario.QtdeCurtidaRespComentario, &comentario.CurtiuRespComentario); err != nil {
			log.Println("Erro ao scanear comentarios do post", err)
			c.Status(400)
			return
		}
		getFullComent = append(getFullComent, comentario)

	}
	c.JSON(200, getFullComent)
}

func RespComment(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)
	var comment = models.RespComment{}

	comment.Resposta = c.PostForm("resposta")
	comment.IDUsuarioRP = IDUser
	comment.IDComentarioRP, _ = strconv.Atoi(c.PostForm("id_comentario"))

	_, err := DB.Exec(`INSERT INTO resp_comentario (resposta, id_comentario_rp, id_usuario_rp)
	VALUES (?, ?, ?)`, comment.Resposta, comment.IDComentarioRP, IDUser)
	if err != nil {
		log.Println("Erro ao responder comentario", err)
		c.Status(400)
	}

	err = DB.QueryRow(`
	SELECT u.nome, u.arroba, u.link_perfil, rc.id_resp_comentario, rc.resposta, rc.dt_resp_comentario 
	FROM resp_comentario AS rc
	INNER JOIN usuario u ON u.id_usuario = rc.id_usuario_rp
	WHERE rc.id_resp_comentario = LAST_INSERT_ID()
	`).Scan(&comment.Nome, &comment.Arroba, &comment.LinkPerfil, &comment.IDResposta, &comment.Resposta, &comment.DTRespCometario)
	if err != nil {
		log.Println("Erro ao scanear tabela resp_comentario", err)
		c.Status(500)
		return
	}

	c.JSON(200, comment)
}

func ListRespComment(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)
	var comentario models.RespComment
	var getFullComent []models.RespComment

	limit, err := strconv.Atoi(c.PostForm("limit"))
	if err != nil {
		log.Println("Erro ao converter limit", err)
		c.Status(400)
		return
	}

	IDRespComment, err := strconv.Atoi(c.PostForm("id_resp_comentario"))
	if err != nil {
		log.Println("Erro ao converter id da resposta comentario", err)
		c.Status(400)
		return
	}

	IDComment, err := strconv.Atoi(c.PostForm("id_comentario"))
	if err != nil {
		log.Println("Erro ao converter id do comentario", err)
		c.Status(400)
		return
	}

	var where = fmt.Sprintf("WHERE rc.id_resp_comentario < %d", IDRespComment)
	if IDRespComment == 0 {
		where = "WHERE 1=1"
	}

	selector := fmt.Sprintf(`SELECT rc.id_resp_comentario, rc.resposta, rc.dt_resp_comentario, rc.id_comentario_rp, u.nome, u.arroba, u.link_perfil,
	COUNT(crc.id_curt_resp_comentario) AS qtde_curtida_resp_comentario,
	(SELECT CASE 
		WHEN id_usuario_curt_resp = %d THEN 1 
		ELSE 0
	END
	FROM curt_resp_comentario 
		WHERE id_curt_resp_comentario = rc.id_resp_comentario
		AND id_usuario_curt_resp = %d
	) AS curtiu_resp_comentario
	FROM resp_comentario rc 
	INNER JOIN	usuario u ON u.id_usuario = rc.id_usuario_rp   
	LEFT JOIN	curt_resp_comentario crc ON crc.id_curt_resp_comentario = rc.id_resp_comentario
	%s
	AND rc.id_comentario_rp = %d
	GROUP BY rc.id_resp_comentario 
	ORDER BY rc.id_resp_comentario DESC
	LIMIT %d`, IDUser, IDUser, where, IDComment, limit)

	query, err := DB.Query(selector)
	if err != nil {
		log.Println("Erro query Get comentario", err)
		c.Status(400)
		return
	}

	for query.Next() {
		if err := query.Scan(&comentario.IDResposta, &comentario.Resposta, &comentario.DTRespCometario, &comentario.IDComentarioRP, &comentario.Nome, &comentario.Arroba, &comentario.LinkPerfil, &comentario.QtdeCurtidadeRespComentario, &comentario.CurtiuRespComentario); err != nil {
			log.Println("Erro ao scanear comentarios do post", err)
			c.Status(400)
			return
		}
		getFullComent = append(getFullComent, comentario)

	}
	c.JSON(200, getFullComent)
}
