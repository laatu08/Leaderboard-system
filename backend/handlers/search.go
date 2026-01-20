package handlers

import (
	"encoding/json"
	"leaderboard/models"
	"leaderboard/store"
	"net/http"
	"sort"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	users := store.SearchUsers(query)

	var response []models.SearchResult
	for _, user := range users {
		response = append(response, models.SearchResult{
			Rank:     store.GetRank(user.Rating),
			Username: user.Username,
			Rating:   user.Rating,
		})
	}

	sort.Slice(response, func(i, j int) bool {
		if response[i].Rank == response[j].Rank {
			// same rank → higher rating first (optional but nice)
			return response[i].Rating > response[j].Rating
		}
		return response[i].Rank < response[j].Rank
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
