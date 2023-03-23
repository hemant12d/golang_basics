package main

import (
	"fmt"
	"log"
	"os"
)

// Custom loggers
var (
	WarningLogger *Logger
	InfoLogger *Logger
	ErrorLogger *Logger
)

func init() {

	// Create/Open a file
	file, err := os.OpenFile("customLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600);
	//fmt.Println(os.O_APPEND, os.O_CREATE, os.O_WRONLY) // 1024 64 1

	// Check for the errors
	if err != nil {
		Fatal("Error occured during file creation", err);
	}

	// We can create custom logger types using the log.New() method
	InfoLogger = New(file, "INFO: ", Ldate|Ltime|Lshortfile)
	WarningLogger = New(file, "WARNING: ", Ldate|Ltime|Lshortfile)
	ErrorLogger = New(file, "ERROR: ", Ldate|Ltime|Lshortfile)
}

func main(){

	fmt.Println("Starting...")
	Println("Server crashed :( ")
	Println("Server got restarted :) ")

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

