package main

import (
	"fmt"
	"log"

	"github.com/abrar-hossain/supplier-risk-radar/config"
	"github.com/abrar-hossain/supplier-risk-radar/internal/company"
	"github.com/abrar-hossain/supplier-risk-radar/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()
	client := company.NewCompanyClient(cfg.CompaniesHouseAPIKey)
	searchHandler := handler.NewSearchHandler(client)
	fmt.Println("Server port:", cfg.ServerPort)
	fmt.Println("API Key loaded:", cfg.CompaniesHouseAPIKey != "")

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/search", searchHandler.Search)

	r.Run(":" + cfg.ServerPort)

}
