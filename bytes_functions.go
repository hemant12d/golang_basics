package main

import (
	"bytes"
	"fmt"
)

func main() {
	 data1 := []byte{102, 97, 108, 99, 111, 110} // falcon
     data2 := []byte{99, 111, 110}                   // on


	 // Bytes Contains function

	// Purpose is to find whether the sub-slice exists within the larger slice
	if bytes.Contains(data1, data2) {
		 fmt.Println("Contains")
	} else {
		 fmt.Println("Does not contain")
	}

	// Equal method of bytes
	if bytes.Equal([]byte("falcon"), []byte("falcon")) {
         fmt.Println("equal")
    } else {
         fmt.Println("not equal")
    }


	// Join methods of bytes
	data3 := [][]byte{[]byte("an"), []byte("old"), []byte("wolf")}
    joined := bytes.Join(data3, []byte(" "))

    fmt.Println(data3)
    fmt.Println(joined)
    fmt.Println(string(joined))


	// Trim method of bytes
	data4 := []byte{32, 32, 102, 97, 108, 99, 111, 110, 32, 32, 32}
	data5 := bytes.Trim(data4, " ")

	fmt.Println(string(data4))
	fmt.Println(data5)
	fmt.Println(string(data5))
}