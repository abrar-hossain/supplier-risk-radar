package main

import (
	"fmt"
	"log"

	"github.com/abrar-hossain/supplier-risk-radar/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()

	fmt.Println("Server port:", cfg.ServerPort)
	fmt.Println("API Key loaded:", cfg.CompaniesHouseAPIKey != "")

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":" + cfg.ServerPort)
}
