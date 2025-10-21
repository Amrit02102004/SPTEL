package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from the LOCAL CI/CD Pipeline! Deployed by Jenkins. A project for SPTEL.\n")
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}