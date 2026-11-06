package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("hello")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env:", err)
	}

	fmt.Println("Loaded PORT:", os.Getenv("PORT"))

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("Port is not found in the environment")
	}

	fmt.Println("PORT : ", portstring)
}
