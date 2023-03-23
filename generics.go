package main

import "fmt"

//Note: This tutorial introduces the basics of generics in Go. With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

// SumFloats adds together the values of m.

// Life without the generics :(
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}


// Life with the generics :)

func SumIntsOrFloats[k comparable, v int64 | float64 ](m map[k]v) v{
	var sum v;
	for _, v := range m {
		sum += v
	}

	return sum;
}
func main() {
	fmt.Println("Starting...");

	intMap := make(map[string]int64)
	intMap["a"] = 1
	intMap["b"] = 2
	intMap["c"] = 3


	floatMap := make(map[string]float64)
	floatMap["a"] = 1.1
	floatMap["b"] = 2.2
	floatMap["c"] = 3.3


	// Without generics
	fmt.Println(SumInts(intMap));
	fmt.Println(SumFloats(floatMap));

	// With generics
	fmt.Println(SumIntsOrFloats(intMap));
	fmt.Println(SumIntsOrFloats(floatMap));


}