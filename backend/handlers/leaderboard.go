package handlers

import (
	"encoding/json"
	"leaderboard/models"
	"leaderboard/store"
	"net/http"
	"strconv"
)

func LeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	limit := 50

	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	var response []models.LeaderboardEntry
	collected := 0

	for rating := 5000; rating >= 100 && collected < limit; rating-- {
		users := store.GetUsersByRating(rating)
		if len(users) == 0 {
			continue
		}

		rank := store.GetRank(rating)

		for _, user := range users {
			if collected >= limit {
				break
			}

			response = append(response, models.LeaderboardEntry{
				Rank:     rank,
				Username: user.Username,
				Rating:   user.Rating,
			})
			collected++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
