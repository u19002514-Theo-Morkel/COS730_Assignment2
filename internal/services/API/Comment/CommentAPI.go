package Comment

import (
	"Assignment2/internal/core/Handlers"
	"Assignment2/internal/core/Models"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ApiController struct {
	*Handlers.HttpClientHandler
}

func (c ApiController) LoadRoutes(h *Handlers.HttpServerHandler) *Handlers.HttpServerHandler {
	slog.Info("Loading routes for comment")

	h.HttpServer.HandleFunc("PUT /comment/", c.CreateComment)
	h.HttpServer.HandleFunc("GET /comment/id/{id}", c.GetAllCommentsForPageID)
	h.HttpServer.HandleFunc("GET /moderation/{id}", c.GetModerationForComment)

	return h
}

func (c ApiController) CreateComment(w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating comment")

	var comment Models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := c.Create(comment, "")

	slog.Debug(resp.Status)

	w.WriteHeader(resp.StatusCode)

}

func (c ApiController) GetAllCommentsForPageID(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get Comments")

	id := r.PathValue("id")

	resp := c.Get("/" + id)

	var comments []Models.Comment
	err := json.NewDecoder(resp.Body).Decode(&comments)
	if err != nil {
		slog.Error(err.Error())
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		return
	}
}

func (c ApiController) GetModerationForComment(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get Moderation for Comment")

	id := r.PathValue("id")

	resp := c.Get("/moderation/" + id)

	var moderation []Models.Moderation
	err := json.NewDecoder(resp.Body).Decode(&moderation)
	if err != nil {
		slog.Error(err.Error())
	}

	err = json.NewEncoder(w).Encode(moderation)
	if err != nil {
		return
	}
}
