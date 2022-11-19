package app

import (
	"database/sql"
	"fmt"
	"github.com/knvovk/copypass/internal/transport/rest"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/knvovk/copypass/internal/config"
	"github.com/knvovk/copypass/internal/service"
	"github.com/knvovk/copypass/internal/storage"

	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, db *sql.DB, log *logrus.Logger) error {
	address := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	router := mux.NewRouter()
	{
		_storage := storage.NewUserStorage(db)
		_service := service.NewUserService(_storage, log)
		_handler := rest.NewUserHandler(_service)
		_handler.Register(router)
	}
	{
		_storage := storage.NewAccountStorage(db)
		_service := service.NewAccountService(_storage, log)
		_handler := rest.NewAccountHandler(_service)
		_handler.Register(router)
	}
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Listen on %s\n", address)
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
