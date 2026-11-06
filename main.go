package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello");

	portstring := os.Getenv("PORT")

	if portstring == ""{
		log.Fatal("Port is not found in the environment")
	}

	fmt.Println("PORT : ",portstring)
}