package Handlers

import (
	"log/slog"
	"net/http"
)

type HttpServerHandler struct {
	HttpServer *http.ServeMux
}

func NewHttpServerHandler(http *http.ServeMux) HttpServerHandler {
	return HttpServerHandler{http}
}

func (h *HttpServerHandler) StartServer(port string) {
	slog.Info("Starting server on port " + port)
	err := http.ListenAndServe(port, h.HttpServer)
	if err != nil {
		panic(err)
	}
}
