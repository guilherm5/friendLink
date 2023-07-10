package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/utils"
)

func Post(c *gin.Context) {
	//Configurando aws para receber post
	service := utils.UtilAWS()
	//Recuperando usuario pelo JWT
	IDJwt := c.GetInt("id")
	legenda := c.PostForm("legenda_post")
	postTexto := c.PostForm("post_texto")

	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.Status(100)
		return
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			log.Println("Error opening file", err)
			c.Status(400)
			return
		}

		_, err = io.ReadAll(src)
		if err != nil {
			log.Println("Error reading file content", err)
			c.Status(400)
			return
		}
		src.Seek(0, io.SeekStart)

		uploader := s3manager.NewUploader(service)

		input := &s3manager.UploadInput{
			Bucket: aws.String("frienlinkfotos"),
			Key:    aws.String("post/" + file.Filename),
			Body:   src,
		}

		_, err = uploader.UploadWithContext(context.Background(), input)
		if err != nil {
			log.Println("Error in upload", err)
		}

	}
	linkPost := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/%s", file.Filename)

	query, err := DB.Query("INSERT INTO post (post_imagem, legenda_post, post_texto, id_usuario_pt) VALUES (?, ?, ?, ?)", linkPost, legenda, postTexto, IDJwt)
	if err != nil {
		log.Println("Error in post query", err)
		c.Status(500)
		return
	}

	for query.Next() {
		err := query.Scan(&linkPost, &legenda, postTexto, &IDJwt)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.Status(500)
			return
		}
	}
	c.JSON(200, gin.H{
		"link_image": linkPost,
	})
}
