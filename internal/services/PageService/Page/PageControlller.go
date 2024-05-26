package Page

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
	slog.Info("Loading routes for page")
	h.HttpServer.HandleFunc("PUT /", c.CreatePage)
	h.HttpServer.HandleFunc("GET /{id}", c.GetPageFromID)
	h.HttpServer.HandleFunc("GET /", c.GetALlPages)
	h.HttpServer.HandleFunc("POST /", c.UpdatePage)

	return h
}

func (c Controller) CreatePage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating page")

	var page Models.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	c.DB.Create(&page)

	slog.Info("Page created")

	// Add page to queue
	slog.Info("Adding page to queue")
	c.AddPageToQueue(&page)

	w.WriteHeader(http.StatusCreated)
}

func (c Controller) GetPageFromID(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get Page from ID")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
		return
	}

	page := Models.Page{
		ID: id,
	}

	c.DB.First(&page)

	slog.Info("Getting Page")

	err = json.NewEncoder(w).Encode(page)
	if err != nil {
		return
	}
}

func (c Controller) GetPageFromIDs(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get Page from ID")

	//	get ids from request body
	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	var pages []Models.Page
	c.DB.Find(&pages, ids)

	slog.Info("Getting Pages")
	slog.Debug("Number of pages: " + strconv.Itoa(len(pages)))

	err = json.NewEncoder(w).Encode(pages)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (c Controller) GetALlPages(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting All Pages")

	var pages []Models.Page
	c.DB.Find(&pages)

	log.Print("Getting all Pages")

	err := json.NewEncoder(w).Encode(pages)
	if err != nil {
		return
	}
}

func (c Controller) UpdatePage(w http.ResponseWriter, r *http.Request) {
	log.Print("Update request received")

	var page Models.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.DB.Save(&page)

	log.Print("Updating Page")

	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeletePage(w http.ResponseWriter, r *http.Request) {
	log.Print("Delete request received")

	var page Models.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.DB.Delete(&page)

	log.Print("Deleting Page")

	w.WriteHeader(http.StatusNoContent)
}
