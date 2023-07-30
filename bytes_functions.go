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


	 // Bytes Equal function
	 if bytes.Equal([]byte("falcon"), []byte("falcon")) {
          fmt.Println("equal")
     } else {
          fmt.Println("not equal")
     }
}