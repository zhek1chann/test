package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Print JSON
	fmt.Println("✅ JSON Callback received:")
	for k, v := range data {
		fmt.Printf("%s: %v\n", k, v)
	}

	// Respond OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("JSON received"))
}

func main() {
	http.HandleFunc("/callback", callbackHandler)

	fmt.Println("Listening on http://0.0.0.0:8080/callback ...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
