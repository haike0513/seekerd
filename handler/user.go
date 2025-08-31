package handler

import (
	"seekerd/models"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user models.User
	models.DB.First(&user)
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	user := models.User{Name: name, ID: uuid.New()}
	models.DB.Create(&user)
	c.JSON(200, user)
}
