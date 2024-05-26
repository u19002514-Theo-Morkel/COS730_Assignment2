package Moderation

import (
	"Assignment2/internal/core/Handlers"
	"Assignment2/internal/core/Models"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type Controller struct {
	*Handlers.DbHandler
	*Handlers.RedisHandler
}

func (c Controller) LoadRoutes(h *Handlers.HttpServerHandler) *Handlers.HttpServerHandler {
	slog.Info("Loading routes for comment")

	h.HttpServer.HandleFunc("GET /moderation/{id}", c.GetModerationForComment)

	return h
}

// Run read moderation queue in a go routine
func (c Controller) ProcessModerationQueue() {
	go c.readModerationQueue()
}

// Get moderation for comment from id
func (c Controller) GetModerationForComment(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get Moderation for Comment")

	commentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
		return
	}

	var moderation []Models.Moderation
	c.DB.Model(&Models.Moderation{}).Where("comment_id = ?", commentID).Find(&moderation)

	err = json.NewEncoder(w).Encode(moderation)
	if err != nil {
		return
	}

}
