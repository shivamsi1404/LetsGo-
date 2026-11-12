# LetsGo

An RSS (Really Simple Syndication) feed aggregator built in Go. It lets users register, add RSS feeds, follow them, and collects the latest posts through a background scraper, all exposed through a JSON REST API.

## Features

- User management with API key authentication
- Add and list RSS feeds
- Follow and unfollow feeds
- Background scraper that fetches new posts from feeds
- PostgreSQL storage with type-safe queries

## Tech Stack

| Purpose | Tool |
| --- | --- |
| Language | Go |
| Database | PostgreSQL |
| Query generation | sqlc |
| Migrations | goose |
| Config | Environment variables (`.env`) |

## Project Structure

```
LetsGo-/
├── internal/            # Generated database code (sqlc)
├── sql/                 # Schema migrations and queries
├── main.go              # Entry point and server setup
├── handler_*.go         # HTTP handlers (users, feeds, feed follows, health, errors)
├── middleware_auth.go   # API key authentication
├── scrapper.go          # Background feed scraper
├── rss.go               # RSS parsing
├── models.go            # Data models
├── json.go              # JSON response helpers
└── sqlc.yaml            # sqlc configuration
```

## Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.20 or newer
- [PostgreSQL](https://www.postgresql.org/download/)
- [goose](https://github.com/pressly/goose) (database migrations)
- [sqlc](https://sqlc.dev/) (only needed if you change the SQL queries)

### Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/shivamsi1404/LetsGo-.git
   cd LetsGo-
   ```

2. **Create a `.env` file** in the project root

   ```env
   PORT=8080
   DB_URL=postgres://user:password@localhost:5432/letsgo?sslmode=disable
   ```

3. **Create the database**

   ```bash
   createdb letsgo
   ```

4. **Run the migrations**

   ```bash
   goose -dir sql/schema postgres "<your DB_URL>" up
   ```

5. **Start the server**

   ```bash
   go run .
   ```

The server will be running at `http://localhost:8080`.

## Usage

1. Create a user to receive an API key.
2. Send the key in the `Authorization` header: `Authorization: ApiKey <your_key>`.
3. Add feeds, follow them, and let the scraper collect posts.

## Future Improvements

- Support for multiple feed types, including Atom
- UI/UX for browsing feeds and posts
- Save specific posts to a personal library
- Deploy

## Author

[shivamsi1404](https://github.com/shivamsi1404)