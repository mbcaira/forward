package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var (
	dbConn *pgxpool.Pool
	dbName string
)

func InitializeDBPool(ctx context.Context) {
	dbUrl := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("could-not-connect-to-db")
	}
	dbConn = dbpool
	dbName = os.Getenv("DATABASE_NAME")
}

func CloseDBPool() {
	dbConn.Close()
}

func QueryUrlEntries(shortUrl string) (result string, urlFound bool, err error) {
	if dbConn != nil {
		query := fmt.Sprintf("SELECT longUrl FROM %v WHERE shortUrl = '%v';", dbName, shortUrl)
		err = dbConn.QueryRow(context.Background(), query).Scan(&result)
		if err != nil {
			log.Info().AnErr("queryError", err).Str("query", query).Msg("queryUrl-query-failed")
		}
		if len(result) > 0 {
			urlFound = true
		}
	} else {
		log.Error().Msg("queryUrlEntries-dbConn-not-initialized")
	}

	return result, urlFound, err
}

func WriteUrlEntry(uuid string, shortUrl string, longUrl string) (err error) {
	if dbConn != nil {
		query := fmt.Sprintf("INSERT INTO %v (id, short, long) VALUES ('%v', '%v', '%v')", dbName, uuid, shortUrl, longUrl)
		_, err = dbConn.Exec(context.Background(), query)
		if err != nil {
			log.Err(err).Str("query", query).Msg("writeUrl-query-failed")
		}
	} else {
		log.Error().Msg("writeUrlEntry-dbConn-not-initialized")
	}

	return err
}
