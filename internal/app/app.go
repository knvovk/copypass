package app

import (
	"database/sql"
	"fmt"
	"github.com/knvovk/copypass/internal/transport/rest"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/knvovk/copypass/internal/services"
	"github.com/knvovk/copypass/internal/storages"
)

func Run(db *sql.DB) {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	address := fmt.Sprintf("%s:%s", host, port)
	router := mux.NewRouter()
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	{
		storage := storages.NewUserStorage(db)
		service := services.NewUserService(storage)
		handler := rest.NewUserHandler(service)
		handler.Register(router)
	}

	{
		storage := storages.NewAccountStorage(db)
		service := services.NewAccountService(storage)
		handler := rest.NewAccountHandler(service)
		handler.Register(router)
	}

	log.Printf("Listen on %s...\n", address)
	log.Fatal(server.ListenAndServe())
}
