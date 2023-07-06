package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/database"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.Init()

func NewUser(c *gin.Context) {
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	senha := c.PostForm("senha")

	//Verificano se email é válido
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("Invalid email", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Invalid email": err.Error(),
		})
		return
	}

	//Fazendo hash do password
	passwordUser, err := bcrypt.GenerateFromPassword([]byte(senha), 14)

	//Preparando insert (resolvi usar para fins de teste, mas poso somente executra direto o insert)
	query, err := DB.Prepare(`INSERT INTO usuario (nome, email, senha) VALUES (?,?,?)`)
	if err != nil {
		log.Println("Error in prepare query insert", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in prepare query insert": err.Error(),
		})
		return
	} else {
		//Executando insert
		query.Exec(nome, email, passwordUser)
		c.JSON(http.StatusOK, "Success")
	}
}

func InfoUser(c *gin.Context) {
	//Recuperando usuario pelo JWT
	IDJwt := c.GetInt("id")

	//Começando update
	bio := c.PostForm("bio")
	arroba := c.PostForm("arroba")
	foto, err := c.FormFile("foto")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error retrieving file": err.Error(),
		})
		return
	}
	var data []byte
	if foto != nil {
		src, err := foto.Open()
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

	fotoCapa, err := c.FormFile("foto_capa")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in retrieving file", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error retrieving file": err.Error(),
		})
		return
	}

	var dataCapa []byte
	if fotoCapa != nil {
		srcCapa, err := fotoCapa.Open()
		if err != nil {
			log.Println("Error opening file", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error opening file": err.Error(),
			})
			return
		}

		dataCapa, err = ioutil.ReadAll(srcCapa)
		if err != nil {
			log.Println("Error reading file content", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error reading file content": err.Error(),
			})
			return
		}
	}

	query, err := DB.Query(`UPDATE usuario SET foto = ?, foto_capa = ?, bio = ?, arroba = ? WHERE id_usuario = ? `, data, dataCapa, bio, arroba, IDJwt)
	if err != nil {
		log.Println("Error in update query", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in update query": err.Error(),
		})
		return
	}

	for query.Next() {
		err := query.Scan(&data, &dataCapa, &bio, &arroba)
		if err != nil {
			log.Println("Error in scan result query", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Error in scan result query": err.Error(),
			})
			return
		}
	}
}
