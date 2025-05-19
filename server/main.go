package main

import (
	"fmt"
	"log"
	"net/http"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Read the body (if needed)
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Log received data
		fmt.Println("Callback received:")
		for key, values := range r.Form {
			for _, value := range values {
				fmt.Printf("%s = %s\n", key, value)
			}
		}

		// Respond
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Callback received"))
	} else {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/callback", callbackHandler)

	// Bind to all interfaces (0.0.0.0), not just localhost
	fmt.Println("Listening on http://0.0.0.0:8081...")
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
