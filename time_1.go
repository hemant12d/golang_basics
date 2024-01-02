package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	now := time.Now()

	// Get the number of milliseconds in 30 days
	var millis time.Duration = 30 * 24 * time.Hour

	// Add 30 days to the current time
	plus30days := now.Add(millis)

	fmt.Println(plus30days.UnixMilli())
}