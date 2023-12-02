package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main(){
	pubKey, priKey, err := ed25519.GenerateKey(rand.Reader)

	fmt.Printf("pubKey: %v\n", priKey.Public().(ed25519.PublicKey))
	fmt.Printf("priKey: %v\n", pubKey)
	fmt.Printf("error: %v\n", err)
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