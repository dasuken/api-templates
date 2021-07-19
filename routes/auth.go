package routes

import (
	"github.com/noguchidaisuke/api-templates/handlers"
	"net/http"
)

func NewAuthRoutes(authHandlers handlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path: "/users",
			Method: http.MethodGet,
			Handler: authHandlers.GetUsers,
		},
	}
}