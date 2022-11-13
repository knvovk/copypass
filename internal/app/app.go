package app

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knvovk/copypass/internal/config"
	"github.com/knvovk/copypass/internal/domain"
	"github.com/knvovk/copypass/internal/handler"
	"github.com/knvovk/copypass/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, pool *pgxpool.Pool, log *logrus.Logger) error {
	e := echo.New()
	{
		r := domain.NewUserRepository(pool)
		s := service.NewUserService(r, log)
		h := handler.NewUserHandler(s)
		h.Register(e)
	}
	{
		r := domain.NewAccountRepository(pool)
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
