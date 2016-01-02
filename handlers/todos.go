package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vallard/trudiedo-api/services"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	ts := services.Todos.List()
	if err := json.NewEncoder(w).Encode(ts); err != nil {
		log.Println("%s", err)
	}
}
