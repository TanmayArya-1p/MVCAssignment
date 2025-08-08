package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupStaticFileRoutes(root *mux.Router) {
	root.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
