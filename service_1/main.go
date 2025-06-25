package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    // Liveness endpoint
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        jsonResponse(w, map[string]string{
            "status":  "ok",
            "service": "1",
        })
    })

    // Health check endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
    })

    // Business endpoint
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        jsonResponse(w, map[string]string{
            "message": "Hello from Service 1",
        })
    })

    addr := ":5001"
    log.Printf("Service 1 listening on port %s…", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

// jsonResponse writes a JSON-encoded response with Content-Type header.
func jsonResponse(w http.ResponseWriter, data map[string]string) {
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, `{"error":"encoding_failed"}`, http.StatusInternalServerError)
    }
}
