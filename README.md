# Music Catalog

**Music Catalog** is a Go-based web application that allows users to manage music preferences, track activity, and interact with the Spotify API. It provides user authentication, track activity management, and Spotify track search.

---

## 🚀 Features

- **User Management**: Register, log in, and manage user accounts.
- **Track Activities**: Like/dislike tracks and save user-specific track activities.
- **Spotify Integration**: Search for tracks using Spotify's API. Get personalized track recommendations.
- **Database Integration**: Uses PostgreSQL for data persistence.
- **RESTful API**: Provides endpoints for client interaction.

---

## 🗂️ Project Structure

```text
cmd/
  └── main.go                  # Application entry point
internal/
  └── configs/                # Configuration management
  └── handler/                # HTTP handlers for API endpoints
  └── models/                 # Data models
  └── repository/             # Database and API repository
  └── service/                # Business logic layer
pkg/
  └── httpclient/             # HTTP client utilities
internalsql/
  └── (files)                 # Database connection utilities
```

---

## ⚙️ Prerequisites

- Go 1.20 or later
- Docker (for running the PostgreSQL database)
- Spotify API credentials (Client ID & Client Secret)

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

   Update the database connection string and Spotify API credentials in `config.yaml`.

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

   The application will run on the port specified in the config file (default: `:9999`).

---

## 📡 API Endpoints

### 🔐 Memberships

- **Sign Up**  
  `POST /memberships/sign_up`

  **Request Body:**
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

  **Request Body:**
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