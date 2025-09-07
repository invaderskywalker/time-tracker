package sources

import (
	"log"
	"time-tracker/backend/config"
	"time-tracker/backend/sources/database"
)

type Sources struct {
	DB *database.SQLite
	// @TODO:: include rest here
}

func InitiateSources() (*Sources, error) {
	db, err := database.NewSQLite(config.SQLITE_PATH)
	if err != nil {
		log.Println("Failed to connect to SQLite:", err)
		return nil, err
	}
	return &Sources{
		DB: db,
	}, nil
}
