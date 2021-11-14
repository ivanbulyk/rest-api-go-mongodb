package user

import (
	"fmt"
	"github.com/ivanbulyk/rest-api-go-mongodb/internal/apperror"
	"github.com/ivanbulyk/rest-api-go-mongodb/internal/handlers"
	"github.com/ivanbulyk/rest-api-go-mongodb/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return  &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost ,usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut,userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return  apperror.ErrNotFound
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error{
	return  apperror.NewAppError(nil, "test", "test", "t213")
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error{
	return fmt.Errorf("this is an API error")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("This is update user"))

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("This is partially update user"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("This is delete user"))

	return nil
}
