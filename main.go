package main

import (
	"fmt"
	"seekerd/cmd"
	"seekerd/handler"
	"seekerd/models"
	"seekerd/seeker"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cmd.Execute()
	// gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	seeker.Hello()
	dsn := "host=localhost user=postgres password=postgresdb dbname=seekerd port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	models.InitDB(db)
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("db connected:", db != nil)
	// route.Use(cors.Default())
	// route.Use(gin.Logger())
	v1 := route.Group("/api/v1")
	v1.GET("/user", handler.GetUser)
	route.Run(":8080")
}
