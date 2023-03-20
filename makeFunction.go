package main;

import "fmt"

func main() {
	mymap := make(map[int]string)
	mymap[2] = "TWO" // assign value
	fmt.Println(mymap[2])
	fmt.Println(mymap)

	myList := make([]string,6, 10)
	myList[0] = "London"
	fmt.Println(myList)
	fmt.Println("My list length", len(myList))

	// A channel is created using chan keyword and it can only transfer data of the same type, different types of data are not allowed to transport from the same channel.
	myChannel := make(chan int, 1)
	myChannel <- 10
	fmt.Println(<-myChannel)
}
