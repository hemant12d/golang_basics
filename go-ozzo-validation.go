package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
)
type Gender string
type User struct {
	Name string
	Email string
	Username string
	Gender Gender
}


const (
	Male Gender = "male"
	Female Gender = "female"
)

func (u *User) validateUser() error{
	// Set of rules for the validate the struct
	result := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Gender, validation.Required, validation.In(Male, Female)),
	)

	// Return validation result
	return result;
}

func main(){
	// User Data
	user := &User{
		Name: "Hemant Soni", Email: "hsonian@gmail.com", Username: "user_name", Gender: "male",
	}

	// Validation function
	result := user.validateUser();

	// Check the validation result
	if err := result; err != nil {
		// Error found
		fmt.Printf("Found error list \n");
		fmt.Println(err);
	}
	if result == nil {
		fmt.Printf("Everything seems working \n")
	}

	// Return the response
}