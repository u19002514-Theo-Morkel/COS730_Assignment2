package main

import (
	"Assignment2/internal/core/Handlers"
	"Assignment2/internal/core/config"
	"Assignment2/internal/services/API/Comment"
	"Assignment2/internal/services/API/Page"
	"Assignment2/internal/services/API/Recommendation"
	"github.com/joho/godotenv"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}
	HttpServer := Handlers.NewHttpServerHandler(config.CreateHTTPServer())

	PageHttpClient := Handlers.NewHttpClientHandler(config.CreateHTTPClient(), "http://localhost:8080")
	pageApi := Page.ApiController{HttpClientHandler: &PageHttpClient}
	pageApi.LoadRoutes(&HttpServer)

	RecHttpClient := Handlers.NewHttpClientHandler(config.CreateHTTPClient(), "http://localhost:7999")
	recApi := Recommendation.ApiController{HttpClientHandler: &RecHttpClient}
	recApi.LoadRoutes(&HttpServer)

	CommentHttpClient := Handlers.NewHttpClientHandler(config.CreateHTTPClient(), "http://localhost:8081")
	commentApi := Comment.ApiController{HttpClientHandler: &CommentHttpClient}
	commentApi.LoadRoutes(&HttpServer)

	HttpServer.StartServer(":8000")
}
