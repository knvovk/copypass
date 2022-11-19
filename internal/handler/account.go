package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/service"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) Register(router *mux.Router) {
	router.HandleFunc("/accounts", h.GetMany).Methods(http.MethodGet)
	router.HandleFunc("/accounts", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}", h.GetOne).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{id}", h.Update).Methods(http.MethodPatch)
	router.HandleFunc("/accounts/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var account data.Account
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}
	defer r.Body.Close()

	_account, err := h.accountService.Create(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data.BuildSuccessResponse(_account))
	w.Write(response)
}

func (h *AccountHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	account, err := h.accountService.GetOne(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data.BuildSuccessResponse(account))
	w.Write(response)
}

func (h *AccountHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	var limit, offset int = 10, 0
	if r.URL.Query().Has("limit") {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}
	if r.URL.Query().Has("offset") {
		offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	}
	w.Header().Set("Content-Type", "application/json")

	accounts, err := h.accountService.GetMany(limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	m := map[string]any{"count": len(accounts), "accounts": accounts}
	response, _ := json.Marshal(data.BuildSuccessResponse(m))
	w.Write(response)
}

func (h *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account := data.Account{Id: params["id"]}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}
	defer r.Body.Close()

	_account, err := h.accountService.Update(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data.BuildSuccessResponse(_account))
	w.Write(response)
}

func (h *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account := data.Account{Id: params["id"]}
	w.Header().Set("Content-Type", "application/json")

	if err := h.accountService.Delete(account); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
