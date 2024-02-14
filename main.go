package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", handleTime)
	serverAddress := ":8795"
	log.Printf("Server running on http://localhost%s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	currentTime := getCurrentTime()
	response := TimeResponse{Time: currentTime}
	respondWithJSON(w, http.StatusOK, response)
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

type TimeResponse struct {
	Time string `json:"time"`
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
