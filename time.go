package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	fmt.Println(time.Time{})
	fmt.Println(time.Now())
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Now().Add(time.Hour))
	fmt.Println(time.Now().Add(-time.Hour))
}
