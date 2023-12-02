package main

import (
	"fmt"
)

var name = "something lorem ipusm doler sit here i am"
func main() {
	// Example byte slice
	byteData := []byte(name)

	// Convert []byte to []int
	intSlice := convertBytesToIntSlice(byteData)
	fmt.Println("Int Slice:", intSlice)

	// Convert []byte to []rune
	runeSlice := convertBytesToRuneSlice(byteData)
	fmt.Println("Rune Slice:", runeSlice)

	// Convert []byte to string
	str := string(byteData)
	fmt.Println("String:", str)
	fmt.Println("String:", string(runeSlice))
}

// Convert []byte to []int
func convertBytesToIntSlice(data []byte) []int {
	intSlice := make([]int, len(data))
	for i, b := range data {
		intSlice[i] = int(b)
	}
	return intSlice
}

// Convert []byte to []rune
func convertBytesToRuneSlice(data []byte) []rune {
	runeSlice := make([]rune, len(data))
	for i, b := range data {
		runeSlice[i] = rune(b)
	}
	return runeSlice
}
