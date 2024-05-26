package Handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

type HttpClientHandler struct {
	HttpClient *http.Client
	Address    string
}

func NewHttpClientHandler(http *http.Client, addr string) HttpClientHandler {
	return HttpClientHandler{http, addr}
}

func (h *HttpClientHandler) Create(body any, path string) *http.Response {
	if path == "" {
		path = "/"
	}

	bodybuf := new(bytes.Buffer)
	err := json.NewEncoder(bodybuf).Encode(body)
	if err != nil {
		return nil
	}

	// Create a new request
	req, err := http.NewRequest("PUT", h.Address+path, bodybuf)
	if err != nil {
		slog.Error(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	slog.Debug("PUT " + h.Address + path)
	// Send the request
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		slog.Error(err.Error())
	}
	return resp
}

func (h *HttpClientHandler) Get(path string) *http.Response {
	if path == "" {
		path = "/"
	}

	slog.Debug("GET " + h.Address + path)

	// Create a new request
	req, err := http.NewRequest("GET", h.Address+path, nil)
	if err != nil {
		slog.Error(err.Error())
	}

	// Send the request
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return resp
}

func (h *HttpClientHandler) Post(body any, path string) *http.Response {
	if path == "" {
		path = "/"
	}

	bodybuf := new(bytes.Buffer)
	err := json.NewEncoder(bodybuf).Encode(body)
	if err != nil {
		return nil
	}
	// Create a new request
	req, err := http.NewRequest("POST", h.Address+path, bodybuf)
	if err != nil {
		slog.Error(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return resp

}

func (h *HttpClientHandler) Delete(path string) *http.Response {
	if path == "" {
		path = "/"
	}

	slog.Debug("DELETE " + h.Address + path)

	// Create a new request
	req, err := http.NewRequest("DELETE", h.Address+path, nil)
	if err != nil {
		slog.Error(err.Error())
	}

	// Send the request
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return resp
}

func (h *HttpClientHandler) Update(body any, path string) *http.Response {
	if path == "" {
		path = "/"
	}

	bodybuf := new(bytes.Buffer)
	err := json.NewEncoder(bodybuf).Encode(body)
	if err != nil {
		return nil
	}
	// Create a new request
	req, err := http.NewRequest("POST", h.Address+path, bodybuf)
	if err != nil {
		slog.Error(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return resp
}
