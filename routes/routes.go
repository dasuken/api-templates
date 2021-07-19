package routes

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/api-templates/middlewares"
	"net/http"
)

// HandlerFuncに必要な要素をstructに
type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		router.HandleFunc(route.Path, middlewares.LogRequests(route.Handler)).Methods(route.Method)
	}
}
