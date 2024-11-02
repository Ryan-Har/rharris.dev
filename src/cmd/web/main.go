package main

import (
	"github.com/Ryan-Har/rharris.dev/src/internal/config"
	"log/slog"
)

func main() {
	config.InitLogger()
	_, err := config.Load()
	if err != nil {
		slog.Error("error loading config", "error", err)
	}

	
}
