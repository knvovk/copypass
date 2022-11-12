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
	e.PATCH("/users/:id", h.Update)
	e.DELETE("/users/:id", h.Delete)
}

func (h *UserHandler) Create(c echo.Context) error {
	user := new(data.User)
	if err := c.Bind(user); err != nil {
		return BadRequestError(c, err)
	}
	_user, err := h.userService.Create(*user)
	if err != nil {
		return InternalServerError(c, err)
	}
	return SuccessResponse(c, _user)
}

func (h *UserHandler) GetOne(c echo.Context) error {
	user, err := h.userService.GetOne(c.Param("id"), true)
	if err != nil {
		return InternalServerError(c, err)
	}
	return SuccessResponse(c, user)
}

func (h *UserHandler) GetMany(c echo.Context) error {
	limit, offset := 10, 0
	query := echo.QueryParamsBinder(c).Int("limit", &limit).Int("offset", &offset)
	if err := query.BindError(); err != nil {
		return BadRequestError(c, err)
	}
	users, err := h.userService.GetMany(limit, offset)
	if err != nil {
		return InternalServerError(c, err)
	}
	resource := echo.Map{
		"count": len(users),
		"users": users,
	}
	return SuccessResponse(c, resource)
}

func (h *UserHandler) Update(c echo.Context) error {
	user := new(data.User)
	if err := c.Bind(user); err != nil {
		return BadRequestError(c, err)
	}
	user.Id = c.Param("id")
	_user, err := h.userService.Update(*user)
	if err != nil {
		return InternalServerError(c, err)
	}
	return SuccessResponse(c, _user)
}

func (h *UserHandler) Delete(c echo.Context) error {
	user := data.User{Id: c.Param("id")}
	if err := h.userService.Delete(user); err != nil {
		return InternalServerError(c, err)
	}
	return c.NoContent(http.StatusNoContent)
}
