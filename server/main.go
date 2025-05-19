package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received callback: %s\n", r.URL.Path)
		fmt.Fprintln(w, "Callback received")
	})
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
