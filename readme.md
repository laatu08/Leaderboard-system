# 🏆 Scalable Leaderboard System

This project implements a **scalable, tie-aware leaderboard system** designed for correctness, concurrency safety, and clear evolution from **10,000 users to millions**.

The system consists of:
- A **Go backend** (Dockerized) with in-memory ranking and live updates
- A **React Native (Expo Web) frontend** for leaderboard and search

This repository is designed to be **easy to run, easy to evaluate, and easy to extend**.

---

## ✨ Features

- Dense ranking with proper tie handling
- Rank-based leaderboard cutoff (Top 3 / 5 / 10 / 50)
- Live score updates via background goroutines
- User search with global rank resolution
- Concurrency-safe backend design
- Clean scalability path to Redis + DB (documented)

---

## 🧱 Tech Stack

**Backend**
- Go (net/http)
- In-memory data structures
- Docker

**Frontend**
- React Native
- Expo Web
- TypeScript

---

## 📋 Prerequisites

Make sure you have the following installed:

- Git
- Docker + Docker Compose
- Node.js **v18 or higher** (recommended: v20)
- npm

Optional:
- Expo Go app (for mobile testing)

---

## 📦 Installation & Setup

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/laatu08/Leaderboard-system
cd leaderboard
```

Expected structure:

```
leaderboard/
├── backend/
├── frontend/
├── docker-compose.yml
└── README.md
```

---

### 2️⃣ Run Backend (Dockerized)

From the project root:

```bash
docker compose up --build
```

Backend will run at:

```
http://localhost:8080
```

Health check:

```
http://localhost:8080/health
```

---

### 3️⃣ Run Frontend (Expo Web)

Open a **new terminal**:

```bash
cd frontend
npm install
npx expo start --web
```

Frontend will be available at:

```
http://localhost:8081
```

---

## 🔗 API Overview

### Get Leaderboard (Rank-Based)

```
GET /leaderboard?rank=K
```

Returns **all users whose rank ≤ K**, preserving ties.

---

### Search User

```
GET /search?query=username
```

Returns matched users with:
- Username
- Rating
- Global rank

---

## 🔄 Live Updates

- Backend runs a background goroutine that updates random user scores every few seconds
- Frontend polls periodically to reflect live ranking changes

---

## 🧠 System Design Highlights

### 10,000 Users
- In-memory user store
- Rating bucket array (bounded ratings)
- Dense ranking without sorting
- RWMutex for concurrency safety

### Scaling to Millions
- Database as source of truth
- Redis Sorted Sets for live ranking
- Async write pipeline
- Stateless API servers

Detailed design rationale is documented in the project.

---

## 🧹 Stopping the Application

To stop services:

```bash
Ctrl + C
```

---

## 🧪 Troubleshooting

**Backend not responding**
- Ensure Docker is running
- Check logs:
  ```bash
  docker compose logs
  ```

**Frontend API errors**
- Ensure backend is running on port 8080
- Ensure no port conflicts

**Expo issues**
- Clear cache:
  ```bash
  npx expo start -c
  ```

---

## 🎥 Demo

A short demo video is included with the submission explaining:
- System design
- Ranking logic
- Live updates
- Scalability approach

---

## ✅ Summary

This project demonstrates:
- Strong system design fundamentals
- Correct and fair leaderboard logic
- Practical concurrency handling
- Clear scalability thinking
- Clean developer experience

---

Thank you for reviewing this project 🙌
