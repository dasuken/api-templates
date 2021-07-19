package routes

import (
	"github.com/noguchidaisuke/api-templates/handlers"
	"net/http"
)

func NewAuthRoutes(authHandlers handlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path: "/signup",
			Method: http.MethodPost,
			Handler: authHandlers.SignUp,
		},
		{
			Path: "/users/{id}",
			Method: http.MethodGet,
			Handler: authHandlers.GetUser,
		},
		{
			Path: "/users",
			Method: http.MethodGet,
			Handler: authHandlers.GetUsers,
		},
		{
			Path: "/users/{id}",
			Method: http.MethodPut,
			Handler: authHandlers.PutUser,
		},

		{
			Path: "/users/{id}",
			Method: http.MethodDelete,
			Handler: authHandlers.DeleteUser,
		},
	}
}