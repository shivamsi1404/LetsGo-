package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/shivamsi1404/LetsGo-/internal/database"
	 _ "github.com/lib/pq"

)

type apiconfig struct{
	DB *database.Queries
}

func main() {
	fmt.Println("hello")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env:", err)
	}

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("Port is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("database not found")
	}

	conn,err := sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal("database not connected",err)
	}

	queries := database.New(conn)

	apicfg := apiconfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz",handlerReadiness)
	v1router.Get("/err",handlerErr)
	v1router.Post("/users",apicfg.handlerCreateUser)
	v1router.Get("/users",apicfg.middlewareAuth(apicfg.handlerGetUser))

	v1router.Post("/feeds",apicfg.middlewareAuth(apicfg.handlerCreateFeed))
	v1router.Get("/feeds",apicfg.handlerGetFeeds)

	v1router.Post("/feedfollows",apicfg.middlewareAuth(apicfg.handlerCreateFeedFollow))
	v1router.Get("/feedfollows",apicfg.middlewareAuth(apicfg.handlerGetFeedFollows))


	router.Mount("/v1",v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("Server starting on port on %v", portstring)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
