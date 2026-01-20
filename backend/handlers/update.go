package handlers

import (
	"leaderboard/store"
	"net/http"
	"strconv"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	count := 10 // default

	if c := r.URL.Query().Get("count"); c != "" {
		if parsed, err := strconv.Atoi(c); err == nil && parsed > 0 {
			count = parsed
		}
	}

	store.UpdateRandomUsers(count)
	w.Write([]byte("ratings updated"))
}
