package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net/http"
	"os"
	"task3/internal/config"
	"task3/internal/handlers"
	"task3/internal/models"
	"task3/internal/route"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.Info(cfg.HttpServer.Port)
	logger.Info(cfg.HttpServer.Host)

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	h := handlers.New(db)

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Project{})

	r := route.Router(h)

	err = http.ListenAndServe(":"+cfg.HttpServer.Port, r)
	if err != nil {
		logger.Error("Error starting server: %v", err)
	}
}
