package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

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
