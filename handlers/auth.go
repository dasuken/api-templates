package handlers

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/api-templates/util"
	"github.com/noguchidaisuke/micro-service-with-mongo/api/restutil"
	"io"
	"net/http"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {}

func NewAuthHandlers() AuthHandlers {
	return &authHandlers{}
}

func (h *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	util.WriteAsJson(w, http.StatusOK, body)
}
func (h *authHandlers) PutUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	util.WriteAsJson(w, http.StatusOK, vars["id"])
}

func (h *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	util.WriteAsJson(w, http.StatusOK, p)
}
func (h *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	util.WriteAsJson(w, http.StatusOK, p)
}
func (h *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

	p := r.URL.Path
	util.WriteAsJson(w, http.StatusOK, p)
}