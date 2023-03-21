package testDir

import (
	"math/rand"
	"time"
)

func Slogan() string {
	// List of slogans
	sloganList := []string{"Mind your own business is a artistry", "Do what you love", "Here's is the keyboard", "Here's the mouse"}

	// Generate random number [0, 2] range
	rand.Seed(time.Now().UnixNano())
	n := 0 + rand.Intn(3-0+1)

	return sloganList[n]
}
