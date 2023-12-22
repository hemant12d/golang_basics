package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(-time.Second) // Add one hour to current time
	fmt.Println("Before rounding:", t) // Print time before rounding

	roundedT := t.Round(time.Second)           // Round to the nearest hour
	fmt.Println("After rounding:", roundedT) // Print time after rounding
}
