package main

import (
	"log/slog"
	"net/http"
	"os"
	"task3/internal/config"
	"task3/internal/route"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	r := route.Router()

	err := http.ListenAndServe(":"+cfg.HttpServer.Port, r)
	if err != nil {
		logger.Error("Error starting server: %v", err)
	}
}
