package main

import (
	"Assignment2/internal/core/Handlers"
	"Assignment2/internal/core/config"
	"Assignment2/internal/services/CommentService/Comment"
	"Assignment2/internal/services/CommentService/Moderation"
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

	commentController := Comment.Controller{DbHandler: &Db, RedisHandler: &Redis}
	commentController.LoadRoutes(&HttpServer)

	moderationController := Moderation.Controller{DbHandler: &Db, RedisHandler: &Redis}
	moderationController.LoadRoutes(&HttpServer)
	moderationController.ProcessModerationQueue()

	HttpServer.StartServer(":8081")

}
