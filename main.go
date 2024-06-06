package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body.", http.StatusInternalServerError)
		return
	}

	fmt.Println("Received webhook request with body:")
	fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/webhook", webhookHandler)
	fmt.Println("Server is listening on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
