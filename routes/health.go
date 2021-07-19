package routes

import (
	"fmt"
	"net/http"
)

func NewHealthRoutes() []*Route {
	return []*Route {
		{
			Path: "/",
			Method: http.MethodGet,
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello World!")
			},
		},
	}
}