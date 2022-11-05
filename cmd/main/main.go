package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knvovk/copypass/internal/app"
	"github.com/knvovk/copypass/internal/config"
	"os"
)

var configPath = flag.String("config", "config.toml", "Path to config.toml")

func main() {
	flag.Parse()

	cfg, err := config.GetConfig(*configPath)
	if err != nil {
		fmt.Printf("Can't load config file: %v\n", err)
		os.Exit(1)
	}

	pool, err := pgxpool.New(context.Background(), cfg.DB.URL)
	if err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	if err = app.Run(cfg, pool); err != nil {
		fmt.Printf("Unable to run application: %v\n", err)
		os.Exit(1)
	}
}
