package main

import (
	"apiTest.com/initializers"
	"apiTest.com/models"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	initializers.DB.AutoMigrate(&models.Test{})
	fmt.Println("? Migration complete")
}
