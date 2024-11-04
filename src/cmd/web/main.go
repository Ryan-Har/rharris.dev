package main

import (
	"github.com/Ryan-Har/rharris.dev/src/internal/config"
	"github.com/Ryan-Har/rharris.dev/src/internal/handlers"
	"github.com/Ryan-Har/rharris.dev/src/internal/middleware"
	"log/slog"
	"net/http"
)

func main() {
	config.InitLogger()
	cfg, err := config.Load()
	if err != nil {
		slog.Error("error loading config", "error", err)
	}

	mux := http.NewServeMux()

	//pages
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /experience", handlers.Experience)

	//statics
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	loggingMux := middleware.Logging(mux)

	slog.Info("Starting http server", "port", cfg.HTTPSettings.Port)
	if err := http.ListenAndServe(":"+cfg.HTTPSettings.Port, loggingMux); err != nil {
		slog.Error(err.Error())
	}
}
