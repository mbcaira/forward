package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func getDBConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error().Msg("could-not-connect-to-db")
	}
	return conn
}

func QueryUrlEntries(longUrl string) (result string, urlFound bool, err error) {
	conn := getDBConnection()
	defer conn.Close(context.Background())
	var matchingLongUrl string
	conn.QueryRow(context.Background(), "")
}

func WriteUrlEntry(uuid string, shortUrl string, longUrl string) (id string, err error) {}

func ReadDb(query string) (result []string, err error) {}

func WriteDb() (id string, err error) {}
