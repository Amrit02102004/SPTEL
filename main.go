package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}


var (
	notesStore = make(map[string]Note)
	storeLock  = sync.RWMutex{}
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		storeLock.RLock()
		defer storeLock.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notesStore)
	case http.MethodPost:
		var note Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		storeLock.Lock()
		defer storeLock.Unlock()

		notesStore[note.Title] = note
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Note created successfully")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/notes", notesHandler)

	log.Println("Notes API server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}