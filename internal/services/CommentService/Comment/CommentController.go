package Comment

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

	h.HttpServer.HandleFunc("PUT /", c.CreateComment)
	h.HttpServer.HandleFunc("GET /{id}", c.GetAllCommentsForPageID)

	return h
}

func (c Controller) CreateComment(w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating comment")

	var comment Models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	c.DB.Create(&comment)

	slog.Info("Comment created")

	// Add comment to queue
	slog.Info("Adding comment to queue")
	c.AddCommentToQueue(&comment)

	w.WriteHeader(http.StatusCreated)
}

func (c Controller) GetAllCommentsForPageID(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get Comments")

	pageID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
		return
	}

	var comments []Models.Comment
	c.DB.Model(&Models.Comment{}).Preload("Page").Where("page_id = ?", pageID).Find(&comments)

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		slog.Error(err.Error())
	}
}
