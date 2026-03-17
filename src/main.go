package main

import (
	"encoding/json"
	"fmt"
	="net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/celestial-bodies", getCelestialBodies).Methods("GET")
	fmt.Println("Server is running on port 8000...")
	http.ListenAndServe(":8000", r)
}

func getCelestialBodies(w http.ResponseWriter, r *http.Request) {
	celestialBodies := []string{ "Sun", "Moon", "Earth" }
	json.NewEncoder(w).Encode(celestialBodies)
}
