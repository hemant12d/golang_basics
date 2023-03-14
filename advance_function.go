package main;

import "fmt"

func main() {
	fmt.Println("Hello World");
	fmt.Printf("%T\n", returnParam)
	fmt.Printf("%T\n", invokeParam(7, returnParam));
}


// First class functions

// Helper function
func returnParam(val int) int {
	return val;
}

func invokeParam(val int, f func(int) int)int{
	return f(val);
};