package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

var message = "abc"

func main(){
	// Key generation
	pubKey, priKey, err := ed25519.GenerateKey(rand.Reader)
	fmt.Printf("pubKey: %v\n", pubKey)

	fmt.Printf("public key as string: %s\n", base64.StdEncoding.EncodeToString(pubKey))

	pubKeyDecode, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(pubKey))
	fmt.Printf("pubKey decoded: %v\n", pubKeyDecode)

	fmt.Printf("priKey: %v\n", base64.StdEncoding.EncodeToString(priKey))
	fmt.Printf("error: %v\n", err)

	// Sign the message using the private key
	signature := ed25519.Sign(priKey, []byte(message))
	strSignature := base64.StdEncoding.EncodeToString(signature)
	fmt.Printf("signature: %s\n", strSignature)

	// Verifying the signature using public key
	decodeSignature, err2 := base64.StdEncoding.DecodeString(strSignature)
	if err2 != nil{
		fmt.Printf("error during decoding signature: %v", err2)
	}

	// Verify the signature using the public key
	valid := ed25519.Verify(pubKey, []byte(message), decodeSignature)

	if valid {
		fmt.Println("Signature is valid.")
	} else {
		fmt.Println("Signature is not valid.")
	}

}




// Additional encryption information
// Encryption/Decryption
// SSL/PKI
// Signature verification and signature generating
// Key Exchange
// SSL Authentication
// File and Email encryption
// Secure encryption transaction


// Note
// Encrypt and signature generation process is different not same
// Hash, signature, encryption, signing, message is different and have own meaning
// Public and Private key has different-2 uses based on scenario and use cases