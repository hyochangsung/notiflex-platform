package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

var counter atomic.Int64

func podName() string {
	if name := os.Getenv("HOSTNAME"); name != "" {
		return name
	}
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("응답 쓰기 실패: %v", err)
	}
}

func main() {
	pod := podName()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /id", func(w http.ResponseWriter, r *http.Request) {
		id := counter.Add(1)
		writeJSON(w, http.StatusOK, map[string]string{
			"id":           strconv.FormatInt(id, 10),
			"generated_by": pod,
		})
	})

	log.Printf("notiflex-api 시작: pod=%s port=8080", pod)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
