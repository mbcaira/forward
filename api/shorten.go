package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Shorten(c *gin.Context) {
	var message = response{
		Message: "ok",
	}
	c.IndentedJSON(http.StatusOK, message)
}
