package main

import "fmt"

type code struct {
	zipCode int64
	pinCode int64
}

type address struct {
	firstLine string
	secondLine string
	code
}

type user struct {
	name string
	address
}

func main(){

	newUser := user{
		name: "Hemant Soni",
		address: address{
			firstLine: "This is first line",
			code: code{
				pinCode: 3434,
				zipCode: 4546,
			},
		},
	}

	fmt.Printf("new user: %+v\n", newUser)
}