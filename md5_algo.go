package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

func main() {

	data1 := []byte("data 1 is here")
	data2 := []byte("data 1 is here")

	data1_hash := md5.Sum(data1)
	data2_hash := md5.Sum(data2)

	fmt.Printf("%x\n", data1_hash)
	fmt.Printf("%x\n", data2_hash)

	// Compare the two hash values
	if bytes.Equal(data1_hash[:], data2_hash[:]) {
		fmt.Println("The hash values are the same.")
	} else {
		fmt.Println("The hash values are different.")
	}
}