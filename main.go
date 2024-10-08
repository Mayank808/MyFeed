package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"webScraper/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type APIConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment variables.")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB URL is not found in the environment variables.")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error occured when connecting to the database: %v\n", err)
	}

	queries := database.New(connection)

	apiConfig := &APIConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"http://*", "https://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerHealth)
	v1Router.Get("/error", handlerError)

	v1Router.Post("/user", apiConfig.handleCreateUser)
	v1Router.Get("/user", apiConfig.authMiddleware(apiConfig.handlerGetUserByAPIKey))

	v1Router.Get("/social-feeds", apiConfig.handlerGetAllSocialFeeds)
	v1Router.Post("/social-feed", apiConfig.authMiddleware(apiConfig.handleCreateSocialFeed))

	v1Router.Post("/social-feed-followed", apiConfig.authMiddleware(apiConfig.handleCreateFeedFollowed))
	v1Router.Get("/social-feed-followed", apiConfig.authMiddleware(apiConfig.handlerGetAllFollowedSocialFeeds))
	v1Router.Delete("/social-feed-followed/{socialFeedFollowedID}", apiConfig.authMiddleware(apiConfig.handlerDeleteFollowedSocialFeed))

	testRouter := chi.NewRouter()
	testRouter.Post("/rss-parsing", handlerTestRSSParsing)

	router.Mount("/v1", v1Router)
	router.Mount("/test", testRouter)

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Server started running on port %s...\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error occured when on port %s: %v\n", port, err)
	}
}
