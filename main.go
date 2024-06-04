package main

import (
	"context"
	"forward/api"
	"forward/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Initialize DB module and main context
	ctx := context.Background()
	db.InitializeDBPool(ctx)
	defer db.CloseDBPool()

	// Initialize API routes
	router := gin.Default()
	router.POST("api/v1/data/shorten", api.Shorten)
	router.GET("/:shortUrl", api.ShortUrl)
	router.Run()
}
