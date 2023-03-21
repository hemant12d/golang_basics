package main

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(20 + <-ch)
}


func myFuncCharChan(mychnl chan string) {
	for v := 0 v < 4 v++ {
		mychnl <- "Developer "
	}
	close(mychnl)
}


func main() {

	// Creating a channel
	//var mychannel chan int
	//fmt.Println("Value of the channel: ", mychannel)
	//fmt.Printf("Type of the channel: %T \n", mychannel)


	// Creating a channel using make() function
	//mychannel1 := make(chan int)
	//fmt.Println("\nValue of the channel1: ", mychannel1)
	//fmt.Printf("Type of the channel1: %T ", mychannel1)


	// Send Operation via operator in channel

	// Creating a channel with integer
	ch := make(chan int)

	// Calling go routine
	go myfunc(ch)
	ch <- 23

	// Creating a channel with string
	charChannel := make(chan string)
	//go myFuncCharChan(charChannel) // calling go routine

	// Anonymous go routine
	go func() {
		charChannel <- "GFG"
		charChannel <- "gfg"
		charChannel <- "Geeks"
		charChannel <- "GeeksforGeeks"
		close(charChannel)
	}()


	// Finding capacity and length of the channel
	fmt.Println("Capacity of the channel is: ", cap(charChannel))
	fmt.Println("Length of the channel is: ", len(charChannel))


	// Iterating over the channel
	for {
		res, ok := <-charChannel
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}




	// Receive Operation via operator in channel

}


// Note: In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel.
//In the channel, the send and receive operation block until another side is not ready by default.