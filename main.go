package main

import (
	"forward/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("api/v1/data/shorten", api.Shorten)
	router.GET("/:shortUrl", api.ShortUrl)

	router.Run()
}
