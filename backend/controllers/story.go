package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/guilherm5/models"
	"github.com/guilherm5/utils"
)

func NewStory(c *gin.Context) {
	var serviceData models.ObjectUserStory
	service := utils.UtilAWS()
	IDUser := utils.GetUserJWT(c)

	strStory, err := uuid.NewV4()
	if err != nil {
		log.Println("Erro ao gerar uuid story", err)
	}

	file, err := c.FormFile("story")
	if err != nil {
		c.Status(400)
		return
	}

	src, err := file.Open()
	if err != nil {
		log.Println("Erro ao abrir arquivo", err)
		c.Status(400)
		return
	}
	_, err = io.ReadAll(src)
	if err != nil {
		log.Println("Erro ao ler arquivo", err)
		c.Status(400)
		return
	}
	src.Seek(0, io.SeekStart)

	uploader := s3manager.NewUploader(service)
	contentType := file.Header.Get("Content-Type")
	input := &s3manager.UploadInput{
		Bucket:             aws.String("frienlinkfotos"),
		Key:                aws.String("story/" + strStory.String()),
		Body:               src,
		ContentType:        &contentType,
		ContentDisposition: aws.String("inline"),
	}

	_, err = uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		log.Println("Erro ao realizar uplaod da imagem no bucket s3", err)
		c.Status(101)
	}

	linkStory := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/story/%s", strStory)

	_, err = DB.Exec(`INSERT INTO story(id_usuario_story,link_story) VALUES (?,?)`, IDUser, linkStory)
	if err != nil {
		c.Status(400)
		log.Println("Erro ao inserir story", err)
		return
	}

	err = DB.QueryRow(`SELECT 
	s.id_story
	, s.id_usuario_story
	, u.id_usuario
	, s.link_story 
	, u.nome 
	, u.link_perfil 
	FROM story s 
	INNER JOIN usuario u ON u.id_usuario = s.id_usuario_story
	WHERE id_story = LAST_INSERT_ID() `).Scan(&serviceData.IDStory, &serviceData.IDUsuarioStory, &serviceData.IDUsuario, &serviceData.LinkStory, &serviceData.Nome, &serviceData.LinkPerfil)
	if err != nil {
		log.Println("Erro ao obter objeto story", err)
		c.Status(500)
		c.Next()
	}

	c.JSON(200, gin.H{
		"id_story":         &serviceData.IDStory,
		"id_usuario_story": &serviceData.IDUsuarioStory,
		"link_story":       &serviceData.LinkStory,
		"id_usuario":       &serviceData.IDUsuario,
		"nome":             &serviceData.Nome,
		"link_perfik":      &serviceData.LinkPerfil,
	})

}

func GetStory(c *gin.Context) {
	var userStory models.ObjectUserStory
	var dataStory []models.ObjectUserStory

	userStory.IDUsuario, _ = strconv.Atoi(c.PostForm("id_usuario"))

	query, err := DB.Query(`SELECT
		s.id_story
		, s.link_story
		, s.dt_criacao
		, s.id_usuario_story
		, u.id_usuario
		, u.nome
		, u.link_perfil
	FROM story s
	INNER JOIN usuario u ON u.id_usuario = s.id_usuario_story
	WHERE u.id_usuario = ?`, userStory.IDUsuario)
	if err != nil {
		log.Println("Erro ao buscar storys do usuario selecionado", err)
		c.Status(101)
	}

	for query.Next() {
		if err := query.Scan(&userStory.IDStory, &userStory.LinkStory, &userStory.DTCriacao, &userStory.IDUsuarioStory, &userStory.IDUsuario, &userStory.Nome, &userStory.LinkPerfil); err != nil {
			log.Println("Erro ao scanear tabela de usuarios", err)
			c.Status(500)
			return
		}
		dataStory = append(dataStory, userStory)
	}
	c.JSON(200, dataStory)

}

func GetStoryUser(c *gin.Context) {
	IDUser := utils.GetUserJWT(c)
	var userStory models.ObjectUserStory
	var dataStory []models.ObjectUserStory

	userStory.IDUsuario, _ = strconv.Atoi(c.PostForm("id_usuario"))

	query, err := DB.Query(`SELECT u.id_usuario, u.nome, u.link_perfil, u.arroba,
    CONCAT('[', GROUP_CONCAT(
        JSON_OBJECT(
            'id_story', s.id_story,
            'link_story', s.link_story,
            'dt_criacao', s.dt_criacao,
            'vizualizou', IFNULL(
                (SELECT CASE
                    WHEN sv.id_usuario_vizualizado = ? THEN 1
                    ELSE 0
                END
                FROM story_vizualizacao sv
                WHERE sv.id_story_vizualizado = s.id_story AND sv.id_usuario_vizualizado = ?), 0)
        )
    ), ']') AS stories
FROM
    usuario u
LEFT JOIN
    story s ON u.id_usuario = s.id_usuario_story
LEFT JOIN
    seguidor seg ON u.id_usuario = seg.id_usuario_seguindo
WHERE
    seg.id_usuario_segue = ?
GROUP BY
    u.id_usuario 
HAVING 
	COUNT(s.id_story) > 0 `, IDUser, IDUser, IDUser)
	if err != nil {
		log.Println("Erro ao buscar storys do usuario selecionado", err)
		c.Status(101)
	}

	for query.Next() {

		if err := query.Scan(&userStory.IDUsuario, &userStory.Nome, &userStory.LinkPerfil, &userStory.Arroba, &userStory.Story); err != nil {
			log.Println("Erro ao scanear tabela de usuarios", err)
			c.Status(500)
			return
		}
		dataStory = append(dataStory, userStory)
	}
	c.JSON(200, dataStory)
}
