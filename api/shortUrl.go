package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func ShortUrl(c *gin.Context) {
	var message = response{
		Message: "ok",
	}
	c.IndentedJSON(http.StatusOK, message)
}
