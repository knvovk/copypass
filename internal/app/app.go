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
	"github.com/knvovk/copypass/internal/service"
	"github.com/knvovk/copypass/internal/storage"
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
		_storage := storage.NewUserStorage(db)
		_service := service.NewUserService(_storage)
		_handler := rest.NewUserHandler(_service)
		_handler.Register(router)
	}

	{
		_storage := storage.NewAccountStorage(db)
		_service := service.NewAccountService(_storage)
		_handler := rest.NewAccountHandler(_service)
		_handler.Register(router)
	}

	log.Printf("Listen on %s...\n", address)
	log.Fatal(server.ListenAndServe())
}
