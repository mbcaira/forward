package main

import (
	"forward/api"
	"forward/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	db.ReadDb()
	router.POST("api/v1/data/shorten", api.Shorten)
	router.GET("/:shortUrl", api.ShortUrl)
	router.Run()
}
