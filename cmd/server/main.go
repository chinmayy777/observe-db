package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ObserveDB is running")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server starting on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed: ", err)
	}
}
