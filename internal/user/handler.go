package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest_api/internal/handlers"
	"rest_api/pkg/logging"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func HandlerImpl(logger *logging.Logger) handlers.Handler {
	return &handler{logger: logger}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	write, err := w.Write([]byte("This is list of users"))
	if err != nil {
		return
	} else {
		fmt.Println(write, err)
	}
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	write, err := w.Write([]byte("Get user by id"))
	if err != nil {
		return
	} else {
		fmt.Println(write, err)
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	write, err := w.Write([]byte("Create user"))
	if err != nil {
		return
	} else {
		fmt.Println(write, err)
	}
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	write, err := w.Write([]byte("Update user"))
	if err != nil {
		return
	} else {
		fmt.Println(write, err)
	}
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	write, err := w.Write([]byte("Partially update user"))
	if err != nil {
		return
	} else {
		fmt.Println(write, err)
	}
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	write, err := w.Write([]byte("Delete user"))
	if err != nil {
		return
	} else {
		fmt.Println(write, err)
	}
}
