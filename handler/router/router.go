package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	healthHandler := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthHandler.ServeHTTP)

	todoService := service.NewTODOService(todoDB)
	todoHandler := handler.NewTODOHandler(todoService)
	mux.HandleFunc("/todos", todoHandler.ServeHTTP)

	// mux.HandleFunc("/do-panic", handler.NewDoPanicHandler().ServeHTTP)
	mux.Handle("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()))
	return mux
}