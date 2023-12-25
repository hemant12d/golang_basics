package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

func main() {

	data1 := []byte("data 1 is here")
	data2 := []byte("data 1 is here")

	data1Hash := md5.Sum(data1)
	data2Hash := md5.Sum(data2)

	fmt.Printf("%x\n", data1Hash)
	fmt.Printf("%x\n", data2Hash)

	// Compare the two hash values
	if bytes.Equal(data1Hash[:], data2Hash[:]) {
		fmt.Println("The hash values are the same.")
	} else {
		fmt.Println("The hash values are different.")
	}
}
