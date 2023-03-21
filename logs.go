package main

import (
	"fmt"
	"log"
	"os"
)

func main(){

	fmt.Println("Starting...")
	log.Println("Server crashed :( ")
	log.Println("Server got restarted :) ")

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Logging has completed")
}

