package routes

import (
	"github.com/gorilla/mux"
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
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
