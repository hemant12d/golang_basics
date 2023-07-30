package main

import (
	"fmt"
	"io/ioutil"
)
func main() {
	ct, err := ioutil.ReadFile("./words.txt")
	if err != nil {
		fmt.Println("Error reading words.txt", err)
	}

	fmt.Println("Content of words.txt in bytes", ct)
	fmt.Println("Content of words.txt in words \n", string(ct))
}

