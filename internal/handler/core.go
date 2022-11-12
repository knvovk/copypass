package handler

import (
	"net/http"

	"github.com/knvovk/copypass/internal/data"
	"github.com/labstack/echo/v4"
)

func BadRequestError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, data.Response{
		Status:  data.StatusFailure,
		Message: err.Error(),
		Data:    nil,
	})
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, data.Response{
		Status:  data.StatusFailure,
		Message: err.Error(),
		Data:    nil,
	})
}

func SuccessResponse(c echo.Context, resource any) error {
	return c.JSON(http.StatusOK, data.Response{
		Status:  data.StatusSuccess,
		Message: "",
		Data:    resource,
	})
}
