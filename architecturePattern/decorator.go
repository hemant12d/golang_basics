package main;

import (
	"fmt"
);


// Alias the function type
//type piType func(int) Float64;


// Wrap logger function
//func wrapLogger(fun piType, logger *log.Logger) piType {
//    return func(n int) float64 {
//		// Store rerurn value in result
//		fn := func(n int) (result float64) {
//			defer func(t time.Time) {
//				logger.Printf("%v took %v n %v result", time.Since(t), n, result)
//			}(time.Now())
//			return fun(n)
//		}
//
//		return fn(n)
//	}
//}

func PI(n int) int {
	// make() allocates memory on the heap and initializes and puts zero or empty strings into values. Unlike new(), make() returns the same type as its argument.
	// make() returns a reference to the map, slice, or channel that is allocated on the memory.
	ch := make(chan float64);
	fmt.Println(ch);
	return 3;
}

// Spawn a new go routine for each of the following iterations

func main(){
	// Calculate the approximatation of PI
	fmt.Println(PI(10000))
	fmt.Println(PI(50000))
}
