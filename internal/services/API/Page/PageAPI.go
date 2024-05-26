package Page

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
	slog.Info("Loading Routes for PageAPI")

	h.HttpServer.HandleFunc("PUT /page/", c.CreatePage)
	h.HttpServer.HandleFunc("GET /page/id/{id}", c.GetPageFromID)
	h.HttpServer.HandleFunc("GET /page/all", c.GetAllPages)

	return h
}

func (c ApiController) CreatePage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating page")

	var page Models.Page

	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		slog.Error("Error decoding page")
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := c.Create(page, "")

	slog.Debug(resp.Status)

	w.WriteHeader(resp.StatusCode)
}

func (c ApiController) GetPageFromID(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting page from ID")

	id := r.PathValue("id")

	resp := c.Get("/" + id)

	var page Models.Page
	err := json.NewDecoder(resp.Body).Decode(&page)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Page retrieved")

	err = json.NewEncoder(w).Encode(page)
	if err != nil {
		return
	}
}

func (c ApiController) GetAllPages(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting all pages")

	resp := c.Get("/")

	var pages []Models.Page
	err := json.NewDecoder(resp.Body).Decode(&pages)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Pages retrieved")

	err = json.NewEncoder(w).Encode(pages)
	if err != nil {
		return
	}

}
