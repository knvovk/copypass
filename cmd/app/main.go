package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/knvovk/copypass/internal/app"
)

func main() {
	flags := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(os.Getenv("LOG_PATH"), flags, 0666)
	if err != nil {
		log.Fatalf("Couldn't open the logs file: %v\n", err)
	}
	log.SetOutput(file)

	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v\n", err)
	}
	defer db.Close()

	app.Run(db)
}
