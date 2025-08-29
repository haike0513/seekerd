package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, Seekerd!")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Seekerd!")
	})
	r.Run()
}
