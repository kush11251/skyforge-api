package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"skyforge-api/src/models"
)

func GetCelestialBody(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	// logic to get celestial body by name
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(name)
}
