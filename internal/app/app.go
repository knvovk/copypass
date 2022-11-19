package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/knvovk/copypass/internal/config"
	"github.com/knvovk/copypass/internal/domain"
	"github.com/knvovk/copypass/internal/handler"
	"github.com/knvovk/copypass/internal/service"

	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, db *sql.DB, log *logrus.Logger) error {
	address := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	router := mux.NewRouter()
	{ // User
		r := domain.NewUserRepository(db)
		s := service.NewUserService(r, log)
		h := handler.NewUserHandler(s)
		h.Register(router)
	}
	{ // Account
		r := domain.NewAccountRepository(db)
		s := service.NewAccountService(r, log)
		h := handler.NewAccountHandler(s)
		h.Register(router)
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
