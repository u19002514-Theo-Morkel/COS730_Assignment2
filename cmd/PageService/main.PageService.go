package main

import (
	"Assignment2/internal/core/Handlers"
	"Assignment2/internal/core/config"
	"Assignment2/internal/services/PageService/Page"
	"github.com/joho/godotenv"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warn message")
	slog.Error("Error message")

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	Db := Handlers.NewDB(config.ConnectDB())
	Redis := Handlers.NewRedis(config.CreateRedis())
	HttpServer := Handlers.NewHttpServerHandler(config.CreateHTTPServer())

	pageController := Page.Controller{DbHandler: &Db, RedisHandler: &Redis}
	pageController.LoadRoutes(&HttpServer)

	HttpServer.StartServer(":8080")
}
