package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/database"
	"github.com/guilherm5/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.Init()

func NewUser(c *gin.Context) {
	var newUser models.User
	fmt.Println(DB)
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Println("Error in read body", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error in read body": err.Error(),
		})
		return
	}

	passwordUser, err := bcrypt.GenerateFromPassword([]byte(newUser.Senha), 14)

	query, err := DB.Prepare(`INSERT INTO usuario (nome, email, senha) VALUES (?,?,?)`)
	if err != nil {
		log.Println("Error in prepare query insert", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error in prepare query insert": err.Error(),
		})
		return
	} else {
		query.Exec(newUser.Nome, newUser.Email, passwordUser)
		log.Println("Success")
		c.JSON(http.StatusOK, "Success")
	}

}

func Hello(c *gin.Context) {
	c.JSON(200, "hello")
	c.JSON(200, "depasdsadloyasdasdment succesasdasdsfull")
}
