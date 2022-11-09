package handler

import (
	"net/http"

	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(e *echo.Echo) {
	e.GET("/users", h.GetMany)
	e.POST("/users", h.Create)
	e.GET("/users/:id", h.GetOne)
	e.PATCH("/users/:id", h.GetOne)
	e.DELETE("/users/:id", h.GetOne)
}

func (h *UserHandler) Create(c echo.Context) error {
	user := new(data.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	_user, err := h.userService.Create(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, data.Response{
		Status:  data.StatusSuccess,
		Message: "",
		Data:    _user,
	})
}

func (h *UserHandler) GetOne(c echo.Context) error {
	user, err := h.userService.GetOne(c.Param("id"), true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, data.Response{
		Status:  data.StatusSuccess,
		Message: "",
		Data:    user,
	})
}

func (h *UserHandler) GetMany(c echo.Context) error {
	limit, offset := 10, 0
	query := echo.QueryParamsBinder(c).Int("limit", &limit).Int("offset", &offset)
	if err := query.BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	users, err := h.userService.GetMany(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, data.Response{
		Status:  data.StatusSuccess,
		Message: "",
		Data: echo.Map{
			"count": len(users),
			"users": users,
		},
	})
}

func (h *UserHandler) Update(c echo.Context) error {
	user := new(data.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	_user, err := h.userService.Update(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, data.Response{
		Status:  data.StatusSuccess,
		Message: "",
		Data:    _user,
	})
}

func (h *UserHandler) Delete(c echo.Context) error {
	user := data.User{Id: c.Param("id")}
	if err := h.userService.Delete(user); err != nil {
		return c.JSON(http.StatusInternalServerError, data.Response{
			Status:  data.StatusFailure,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusNoContent, nil)
}
