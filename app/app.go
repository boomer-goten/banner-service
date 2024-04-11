package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	HTTPServer *http.Server
	log        *slog.Logger
}

func New(log *slog.Logger, router *mux.Router) *App {
	timeout, _ := time.ParseDuration(os.Getenv("SERVER_TIMEOUT"))
	idleTimeout, _ := time.ParseDuration(os.Getenv("SERVER_IDLECON"))
	server := &http.Server{
		Addr:         os.Getenv("HOST"),
		IdleTimeout:  idleTimeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      router,
	}
	return &App{HTTPServer: server, log: log}
}

func (a *App) Run() {
	a.log.Info("Running http server", slog.String("address", a.HTTPServer.Addr))
	a.HTTPServer.ListenAndServe()
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.HTTPServer.Handler.ServeHTTP(w, r)
}

func (a *App) Stop() {
	if err := a.HTTPServer.Shutdown(context.Background()); err != nil {
		a.log.Info("Failed to gracefull shutdown server", slog.String("address", a.HTTPServer.Addr))
	} else {
		a.log.Info("Gracefull stopping server", slog.String("address", a.HTTPServer.Addr))
	}
}
