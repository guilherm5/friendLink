package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	//Recuperando usuario pelo JWT
	IDJwt := c.GetInt("id")
	legenda := c.PostForm("legenda_post")
	postTexto := c.PostForm("post_texto")

	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error retrieving file": err.Error(),
		})
		return
	}

	var data []byte
	if file != nil {
		src, err := file.Open()
		if err != nil {
			log.Println("Error opening file", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error opening file": err.Error(),
			})
			return
		}

		data, err = ioutil.ReadAll(src)
		if err != nil {
			log.Println("Error reading file content", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error reading file content": err.Error(),
			})
			return
		}
	}

	query, err := DB.Query("INSERT INTO post (post_imagem, legenda_post, post_texto, id_usuario_pt) VALUES (?, ?, ?, ?)", data, legenda, postTexto, IDJwt)
	if err != nil {
		log.Println("Error in post query", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in post query": err.Error(),
		})
		return
	}

	for query.Next() {
		err := query.Scan(&data, &legenda, postTexto, &IDJwt)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Error in scan result query": err.Error(),
			})
			return
		}
	}
}
