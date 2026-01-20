package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running okay"))
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
