package main

import (
	"fmt"
	"regexp"
)

func main() {
	//usernameRegex := regexp.MustCompile(`^[[:alpha:][:digit:][:punct:][:space:]]+$`)
	usernameRegex := regexp.MustCompile(`^[\p{L}\p{N}\p{Mn}\p{Mc}\p{Nd}\p{Pc}\p{Cf}\s]+$`)
	inputString := "हेमंत^%$"
	if usernameRegex.MatchString(inputString) {
		fmt.Println("Input is valid")
	} else {
		fmt.Println("Input is not valid")
	}

}
