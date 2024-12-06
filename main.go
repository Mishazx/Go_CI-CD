package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/", HelloHandler)
    log.Printf("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "Hello, Docker CI/CD!",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
