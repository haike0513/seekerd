package main

import (
	"fmt"
	"seekerd/seeker"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println(seeker.Hello())
	r := gin.Default()
	seeker.Hello()
	dsn := "host=localhost user=postgres password=postgresdb dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("db connected:", db != nil)
	r.GET("/", func(c *gin.Context) {
		c.String(200, seeker.Hello())
	})
	r.Run()
}
