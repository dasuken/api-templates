package routes

import (
	"github.com/gorilla/handlers"
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

func WithCORS(router *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Accept", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	// createCorsHandler := handlers.CORS(headers, origins, methods)
	// 引数のoptionをheaderに持つ、serveHTTPメンバ関数を持ったcors構造体を作るメソッドを返す
	// return createCorsHandler(router)
	// 引数にserveHttpをもつHandler interface(max)をとる
	return handlers.CORS(headers, origins, methods)(router)
}