package Recommendation

import (
	"Assignment2/internal/core/Handlers"
	"Assignment2/internal/core/Models"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type ApiController struct {
	*Handlers.HttpClientHandler
}

func (c ApiController) LoadRoutes(h *Handlers.HttpServerHandler) *Handlers.HttpServerHandler {
	slog.Info("Loading Routes for RecAPI")
	h.HttpServer.HandleFunc("GET /search", c.SearchPages)

	return h
}

func (c ApiController) SearchPages(w http.ResponseWriter, r *http.Request) {
	slog.Info("Searching pages")

	fmt.Println("GET params were:", r.URL.Query())
	text := r.URL.Query().Get("text")

	slog.Debug("Query: " + text)

	body := map[string]string{
		"text": text,
	}
	resp := c.Post(body, "/search")
	slog.Debug(resp.Status)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error in search", http.StatusBadRequest)
		return
	}

	var recs []Models.Recommendation
	err := json.NewDecoder(resp.Body).Decode(&recs)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Error in search", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		return
	}
}
