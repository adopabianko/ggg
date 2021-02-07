package main

import (
	"fmt"

	"github.com/adopabianko/ggg-boilerplate/config"
	"github.com/adopabianko/ggg-boilerplate/models"
	"github.com/adopabianko/ggg-boilerplate/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DBUrl(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Book{})

	r := routes.SetupRouter()

	r.Run()
}
