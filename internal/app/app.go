package app

import (
	"database/sql"
	"fmt"

	"github.com/knvovk/copypass/internal/config"
	"github.com/knvovk/copypass/internal/domain"
	"github.com/knvovk/copypass/internal/handler"
	"github.com/knvovk/copypass/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, db *sql.DB, log *logrus.Logger) error {
	e := echo.New()
	{
		r := domain.NewUserRepository(db)
		s := service.NewUserService(r, log)
		h := handler.NewUserHandler(s)
		h.Register(e)
	}
	{
		r := domain.NewAccountRepository(db)
		s := service.NewAccountService(r, log)
		h := handler.NewAccountHandler(s)
		h.Register(e)
	}
	address := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	if err := e.Start(address); err != nil {
		return err
	}
	return nil
}
