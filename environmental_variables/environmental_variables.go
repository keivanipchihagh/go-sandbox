package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	// Write to .env file
	env, _ := godotenv.Unmarshal("KEY=123456789")
	err := godotenv.Write(env, "../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read from .env file
	_ = godotenv.Load("../.env")
	key := os.Getenv("KEY")

	fmt.Println(key)
}
