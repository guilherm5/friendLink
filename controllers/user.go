package controllers

import (
	"log"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/database"
	"github.com/guilherm5/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.Init()

func NewUser(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Println("Error in read body", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error in read body": err.Error(),
		})
		return
	}

	//VERIFICATION EMAIL IS EMAIL
	_, err = mail.ParseAddress(newUser.Email)
	if err != nil {
		log.Println("Invalid email", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Invalid email": err.Error(),
		})
		return
	}

	//HASH PASSWORD
	passwordUser, err := bcrypt.GenerateFromPassword([]byte(newUser.Senha), 14)

	//PREPARE QUERY INSERT
	query, err := DB.Prepare(`INSERT INTO usuario (nome, email, senha) VALUES (?,?,?)`)
	if err != nil {
		log.Println("Error in prepare query insert", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in prepare query insert": err.Error(),
		})
		return
	} else {
		//EXEC INSERT
		query.Exec(newUser.Nome, newUser.Email, passwordUser)
		log.Println("Success")
		c.JSON(http.StatusOK, "Success")
	}

}
