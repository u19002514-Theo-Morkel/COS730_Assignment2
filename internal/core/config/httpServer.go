package config

import (
	"net/http"
)

func CreateHTTPServer() *http.ServeMux {
	router := http.NewServeMux()
	return router
}
