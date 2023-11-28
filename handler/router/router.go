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
	mux.Handle("/os-name", middleware.PutOsNameOnContext(middleware.Recovery(handler.NewDoPanicHandler())))
	mux.Handle("/info", middleware.PutOsNameOnContext(middleware.RequestInfo(middleware.Recovery(handler.NewDoPanicHandler()))))
	mux.Handle("/need-auth", middleware.BasicAuth(middleware.Recovery(handler.NewDoPanicHandler())))
	mux.Handle("/wait5", handler.NewWait5SecHandler())

	return mux
}
