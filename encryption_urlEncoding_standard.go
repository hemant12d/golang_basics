package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Here's the `string` we'll encode/decode.
	data := "abc123!?$*&()'-=@~"

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
