//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// name is a string to identify the function call
//// limit the number of numbers the function will print
//// sleep is the number of seconds before the function prints the next value
//func randSleep(name string, limit int, sleep int) {
//	for i := 1; i <= limit; i++ {
//		fmt.Println(name, rand.Intn(i))
//		time.Sleep(time.Duration(sleep * int(time.Second)))
//	}
//
//	time.Sleep(time.Duration(sleep * int(time.Second)));
//	fmt.Println(name, ": Done");
//}
//
//func main() {
//	go randSleep("first: ", 4, 3)
//	randSleep("second: ", 4, 3)
//}


// Wait group
//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//// wg is the pointer to a waitgroup
//// name is a string to identify the function call
//// limit the number of numbers the function will print
//// sleep is the number of seconds before the function prints the next value
//func randSleep(wg *sync.WaitGroup, name string, limit int, sleep int) {
//	defer wg.Done()
//	for i := 1; i <= limit; i++ {
//		fmt.Println(name, rand.Intn(i))
//		time.Sleep(time.Duration(sleep * int(time.Second)))
//
//	}
//
//}
//func main() {
//	wg := new(sync.WaitGroup)
//
//	//wg.Add(2) informs WaitGroup that it must wait for two goroutines.
//	wg.Add(2)
//	go randSleep(wg, "first:", 10, 2)
//	go randSleep(wg, "second:", 3, 2)
//
//	//wg.Wait() then blocks the execution until the goroutines’ execution completes.
//	wg.Wait()
//	fmt.Println("Cheems community...")
//
//}


// Channel communication
package main

import (
	"fmt"
	"sync"
)

func writeChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		limitchannel <- i
	}

}

func readChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		data := <-limitchannel
		fmt.Println(data)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	limitchannel := make(chan int)
	defer close(limitchannel)
	go writeChannel(wg, limitchannel, 3)
	go readChannel(wg, limitchannel, 3)
	wg.Wait()

}