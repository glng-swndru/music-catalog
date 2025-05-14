# 🎧 Music Catalog

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**Music Catalog** is a Go-based web API that lets users manage music preferences, track activities, and interact with Spotify. It supports user authentication, personalized recommendations, and integrates with PostgreSQL and Spotify APIs.

---

## ✨ Features

- 🔐 **User Authentication** – Sign up & log in using JWT-based tokens.
- 🎵 **Spotify Integration** – Search and get personalized track recommendations.
- ❤️ **Track Activities** – Like/dislike tracks and store user interactions.
- 🗄️ **PostgreSQL Support** – Store users, track activity, and tokens.
- 🌐 **RESTful API** – Clean and modular API architecture.

---

## 📁 Project Structure

```text
├── cmd/                            # Application entry point
├── internal/
│   ├── configs/                    # Application configuration (YAML)
│   ├── handler/                    # HTTP handlers (controllers)
│   │   ├── memberships/            # Handlers for user registration and login
│   │   └── tracks/                 # Handlers for track features
│   ├── middleware/                 # Middleware (auth, logging, etc.)
│   ├── models/                     # Domain models
│   │   ├── memberships/            # User models
│   │   ├── spotify/                # Spotify track models
│   │   └── trackactivities/        # User activity models
│   ├── repository/                 # Data access logic
│   │   ├── memberships/            # DB operations for users
│   │   ├── spotify/                # External API calls to Spotify
│   │   └── trackactivities/        # DB operations for activities
│   └── service/                    # Business logic layer
│       ├── memberships/            # Membership logic
│       └── tracks/                 # Track handling logic
└── pkg/
    ├── httpclient/                 # HTTP client for external services
    ├── internalsql/                # PostgreSQL connection utilities
    └── jwt/                        # JWT generation and validation
```

---

## ⚙️ Prerequisites

- Go 1.20 or later
- Docker (for running the PostgreSQL database)
- Spotify Developer Account (Client ID & Secret)

---

## 🛠️ Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/glng-swndru/music-catalog.git
   cd music-catalog
   ```

2. **Copy and update the configuration file:**

   ```bash
   cp internal/configs/config.example.yaml internal/configs/config.yaml
   ```

   Edit `config.yaml` with your DB credentials and Spotify API keys.

3. **Start PostgreSQL with Docker:**

   ```bash
   docker-compose up -d
   ```

4. **Install dependencies:**

   ```bash
   go mod tidy
   ```

5. **Run the application:**

   ```bash
   make run
   ```

   App runs on `localhost:9999` (or the port defined in config).

---

## 📡 API Endpoints

### 🔐 Memberships

- **Sign Up**  
  `POST /memberships/sign_up`

  **Request:**

  ```json
  {
    "email": "user@example.com",
    "username": "username",
    "password": "password123"
  }
  ```

  **Response:**

  ```json
  {
    "accessToken": "jwt-token"
  }
  ```

- **Login**  
  `POST /memberships/login`

  **Request:**

  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

---

### 🎵 Tracks

- **Search Tracks**  
  `GET /tracks/search`

  **Query Parameters:**

  - `query`: Search keyword (e.g., "Bohemian Rhapsody")
  - `pageSize`: Number of results per page (default: 10)
  - `pageIndex`: Page number (default: 1)

  **Response:** List of tracks with details.

- **Upsert Track Activity**  
  `POST /tracks/track-activity`

  **Request Body:**

  ```json
  {
    "spotifyID": "track-id",
    "isLiked": true
  }
  ```

  **Response:** `200 OK` on success.

---

## ✅ Running Tests

To run all unit tests:

```bash
go test ./...
```
