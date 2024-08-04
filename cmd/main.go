package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net/http"
	"os"
	"task3/internal/config"
	"task3/internal/data"
	"task3/internal/handlers"
	"task3/internal/route"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.Info(cfg.HttpServer.Port)
	logger.Info(cfg.HttpServer.Host)

	dsn := "host=localhost user=postgres password=postgres dbname=p_management port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	h := handlers.New(db)

	db.AutoMigrate(&data.User{}, &data.Task{}, &data.Project{})

	r := route.Router(h)

	err = http.ListenAndServe(":"+cfg.HttpServer.Port, r)
	if err != nil {
		logger.Error("Error starting server: %v", err)
	}
}
