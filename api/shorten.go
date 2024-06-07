package api

import (
	"forward/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ShortenRequestBody struct {
	LongUrl string `json:"longUrl"`
}

func Shorten(c *gin.Context) {
	var body ShortenRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, _, err := db.QueryUrlEntries(body.LongUrl)

	if err != nil {
		log.Err(err).Msg("oops")
	}

	log.Info().Str("longUrl", body.LongUrl).Str("res", res).Msg("received")
}
