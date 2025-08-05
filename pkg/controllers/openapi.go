package controllers

import (
	_ "embed"
	"net/http"
)

func OpenAPISpec(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "openapi.json")
}
