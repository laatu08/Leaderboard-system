package main

import (
	"fmt"
	"leaderboard/store"
	"net/http"
	"leaderboard/handlers"
)

func main() {
	store.InitUsers(10000)
	// store.ShowUsers()

	http.HandleFunc("/leaderboard", handlers.LeaderboardHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running okay"))
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
