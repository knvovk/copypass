package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(router *mux.Router) {
	router.HandleFunc("/users", h.GetMany).Methods(http.MethodGet)
	router.HandleFunc("/users", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", h.GetOne).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", h.Update).Methods(http.MethodPatch)
	router.HandleFunc("/users/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user data.User
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}
	defer r.Body.Close()

	_user, err := h.userService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data.BuildSuccessResponse(_user))
	w.Write(response)
}

func (h *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	user, err := h.userService.GetOne(params["id"], true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data.BuildSuccessResponse(user))
	w.Write(response)
}

func (h *UserHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	var limit, offset = 10, 0
	if r.URL.Query().Has("limit") {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}
	if r.URL.Query().Has("offset") {
		offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	}
	w.Header().Set("Content-Type", "application/json")

	users, err := h.userService.GetMany(limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	m := map[string]any{"count": len(users), "users": users}
	response, _ := json.Marshal(data.BuildSuccessResponse(m))
	w.Write(response)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := data.User{Id: params["id"]}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}
	defer r.Body.Close()

	_user, err := h.userService.Update(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data.BuildSuccessResponse(_user))
	w.Write(response)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := data.User{Id: params["id"]}
	w.Header().Set("Content-Type", "application/json")

	if err := h.userService.Delete(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(data.BuildFailureResponse(err))
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
