package handlers

import (
	"github.com/noguchidaisuke/api-templates/util"
	"net/http"
)

type AuthHandlers interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {}

func NewAuthHandlers() AuthHandlers {
	return &authHandlers{}
}

func (h *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	util.WriteAsJson(w, http.StatusOK, p)
}