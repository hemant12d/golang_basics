package main

import (
	"fmt"
	"log"
	"os"
)

// Custom loggers
var (
	WarningLogger *log.Logger
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
)

func init() {

	fmt.Println("Is init automatically invoked ? :,")
	// Create/Open a file
	file, err := os.OpenFile("customLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600);
	//fmt.Println(os.O_APPEND, os.O_CREATE, os.O_WRONLY) // 1024 64 1

	// Check for the errors
	if err != nil {
		log.Fatal("Error occured during file creation", err);
	}

	// We can create custom logger types using the log.New() method
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main(){

	fmt.Println("Starting...")
	log.Println("Server crashed :( ")
	log.Println("Server got restarted :) ")

	// If the file doesn't exist, create it or append to the file
	//file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.SetOutput(file)
	//log.Println("Logging has completed")


	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	ErrorLogger.Println("Something went wrong")
}

