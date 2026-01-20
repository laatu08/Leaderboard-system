package main

import (
	"fmt"
	"leaderboard/handlers"
	"leaderboard/store"
	"net/http"
	"time"
)

func main() {
	store.InitUsers(10000)
	// store.ShowUsers()
	go func() {
		for {
			store.UpdateRandomUsers(5000)
			time.Sleep(2 * time.Second)
		}
	}()

	http.HandleFunc("/leaderboard", handlers.LeaderboardHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	http.HandleFunc("/simulate-update", handlers.UpdateHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running okay"))
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
