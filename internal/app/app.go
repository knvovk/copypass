package app

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knvovk/copypass/internal/config"
	"github.com/knvovk/copypass/internal/handler"
	"github.com/knvovk/copypass/internal/repository"
	"github.com/knvovk/copypass/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, pool *pgxpool.Pool, log *logrus.Logger) error {
	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo, log)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()

	userHandler.Register(e)
	address := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	if err := e.Start(address); err != nil {
		return err
	}
	return nil
}
