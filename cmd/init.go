package cmd

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func Execute() {
	// Your code here
	fmt.Println("Executing command...")
}

func Cors() {
	// Your CORS implementation here
}
