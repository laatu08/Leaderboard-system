package store

import (
	"fmt"
	"leaderboard/models"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	users         = make(map[int]*models.User)
	ratingBuckets [5001]int
	mu            sync.RWMutex
)

// seeding process
func InitUsers(count int) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= count; i++ {
		rating := rand.Intn(4901) + 100 // to start from 100
		user := &models.User{
			ID:       i,
			Username: "user_" + strconv.Itoa(i),
			Rating:   rating,
		}

		users[i] = user
		ratingBuckets[rating]++
	}
	fmt.Println("Seeding of",count,"user done")
}

func ShowUsers() {
	mu.RLock()
	defer mu.RUnlock()

	for _, user := range users {
		fmt.Printf(
			"ID: %d | Username: %s | Rating: %d\n",
			user.ID,
			user.Username,
			user.Rating,
		)
	}
}

func GetUser(id int) *models.User {
	mu.RLock()
	defer mu.RUnlock()
	return users[id]
}

func GetAllUsers() map[int]*models.User {
	mu.RLock()
	defer mu.RUnlock()
	return users
}

// func GetRank(rating int) int {
// 	mu.RLock()
// 	defer mu.RUnlock()

// 	rank := 1
// 	for r := rating + 1; r <= 5000; r++ {
// 		rank += ratingBuckets[r]
// 	}
// 	return rank
// }

func GetRank(rating int) int {
	mu.RLock()
	defer mu.RUnlock()

	rank := 1
	for r := rating + 1; r <= 5000; r++ {
		if ratingBuckets[r] > 0 {
			rank++
		}
	}
	return rank
}



func GetUsersByRating(rating int) []*models.User {
	mu.RLock()
	defer mu.RUnlock()

	var result []*models.User
	for _, user := range users {
		if user.Rating == rating {
			result = append(result, user)
		}
	}
	return result
}
