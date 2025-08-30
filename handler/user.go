package handler

import (
	"seekerd/models"

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
