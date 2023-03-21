package main

import (
	"fmt"
	"go_basics/testDir" // Import local package
)

func main() {
    content := testDir.SayJSR()
	fmt.Println(content)

	slogan := testDir.Slogan()
	fmt.Println(slogan)
}