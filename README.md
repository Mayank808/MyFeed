# README

## Overview

This project is a web server written in Go. The server provides several endpoints to combine all of the users feeds including news, podcasts, many more to be supported in one place. Leverages RSS to pull new content based on the links the users follows.

## Prerequisites

- Go (version 1.15+)
- PostgreSQL Server setup 
- `godotenv` package
- `chi` package
- `cors` package

## Getting Started

### Installation

1. **Clone the repository:**

    ```sh
    git clone repo
    cd your-repo-name
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Create a `.env` file using the template file:**

    ```sh
    touch .env
    ```

4. **Add the respective environment variables to the `.env` file:**

### Running the Server

1. **Load environment variables and start the server:**

    ```sh
    go run main.go
    ```

2. **Check the console output for the server start message:**

    ```
    Server started running on port your_port...
    ```

## Project Structure

- **main.go**: The main file containing the server setup and route definitions.

## Endpoints

#### Health and Error Check
- `GET /v1/health`: Health check endpoint.
- `GET /v1/error`: Error simulation endpoint.

#### User Management
- `POST /v1/user`: Create a new user.
- `GET /v1/user`: Get user details by API key.

#### Social Feeds
- `GET /v1/social-feeds`: Get all social feeds.
- `POST /v1/social-feed`: Create a new social feed.
- `POST /v1/social-feed-followed`: Follow a social feed.
- `GET /v1/social-feed-followed`: Get all followed social feeds.
- `DELETE /v1/social-feed-followed/{socialFeedFollowedID}`: Unfollow a social feed.

#### Testing
- `POST /test/rss-parsing`: Test RSS parsing functionality.

## Running Tests

To run tests, use the following command:

```sh
go test ./...
```