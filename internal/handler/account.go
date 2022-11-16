package handler

import (
	"net/http"

	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/service"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) Register(e *echo.Echo) {
	e.GET("/accounts", h.GetMany)
	e.POST("/accounts", h.Create)
	e.GET("/accounts/:id", h.GetOne)
	e.PATCH("/accounts/:id", h.Update)
	e.DELETE("/accounts/:id", h.Delete)
}

func (h *AccountHandler) Create(c echo.Context) error {
	account := new(data.Account)
	if err := c.Bind(account); err != nil {
		return BadRequestError(c, err)
	}
	_account, err := h.accountService.Create(*account)
	if err != nil {
		return InternalServerError(c, err)
	}
	return SuccessResponse(c, _account)
}

func (h *AccountHandler) GetOne(c echo.Context) error {
	account, err := h.accountService.GetOne(c.Param("id"))
	if err != nil {
		return InternalServerError(c, err)
	}
	return SuccessResponse(c, account)
}

func (h *AccountHandler) GetMany(c echo.Context) error {
	var limit, offset uint = 10, 0
	query := echo.QueryParamsBinder(c).Uint("limit", &limit).Uint("offset", &offset)
	if err := query.BindError(); err != nil {
		return BadRequestError(c, err)
	}
	accounts, err := h.accountService.GetMany(limit, offset)
	if err != nil {
		return InternalServerError(c, err)
	}
	resource := echo.Map{
		"count":    len(accounts),
		"accounts": accounts,
	}
	return SuccessResponse(c, resource)
}

func (h *AccountHandler) Update(c echo.Context) error {
	account := new(data.Account)
	if err := c.Bind(account); err != nil {
		return BadRequestError(c, err)
	}
	account.Id = c.Param("id")
	_account, err := h.accountService.Update(*account)
	if err != nil {
		return InternalServerError(c, err)
	}
	return SuccessResponse(c, _account)
}

func (h *AccountHandler) Delete(c echo.Context) error {
	account := data.Account{Id: c.Param("id")}
	if err := h.accountService.Delete(account); err != nil {
		return InternalServerError(c, err)
	}
	return c.NoContent(http.StatusNoContent)
}
