package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/knvovk/copypass/internal/app"
	"github.com/knvovk/copypass/internal/config"
	"github.com/sirupsen/logrus"
)

var configPath = flag.String("config", "configs/config.toml", "Path to config.toml")

func main() {
	flag.Parse()

	cfg, err := config.GetConfig(*configPath)
	if err != nil {
		fmt.Printf("Can't load config file: %v\n", err)
		os.Exit(1)
	}

	flags := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(cfg.Log.Path, flags, 0666)
	if err != nil {
		log.Fatalf("Couldn't open the logs file: %v\n", err)
	}
	log.SetOutput(file)

	db, err := sql.Open("pgx", cfg.DB.URL)
	if err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = app.Run(cfg, db); err != nil {
		log.Fatalf("Couldn't run application: %v\n", err)
	}
}

func buildLogFormatter(cfg *config.Config) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat:   cfg.Log.DateTimeFormat,
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	}
}
