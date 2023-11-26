package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil{
		fmt.Printf("Error during getting private key: %v", err)
		return nil, nil
	}

	//fmt.Printf("Private key: %v \n", *privateKey)
	//fmt.Printf("Public key: %v \n", privateKey.PublicKey)

	return privateKey, &privateKey.PublicKey
}

func signMessage(key *rsa.PrivateKey) (string, []byte) {

	info := "most of the things can fix with the time"

	hashed := sha256.Sum256([]byte(info))
	fmt.Printf("generated hash value: %v \n", hashed)

	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Printf("error during generating signature: %v \n", err)
	}
	fmt.Printf("generated signature value: %v \n", signature)

	return info, signature
}

func verifySignature(message string, signature []byte, publicKey *rsa.PublicKey) bool {
	hashed := sha256.Sum256([]byte(message))

	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil{
		fmt.Printf("error during signature verification : %v", err)
		return false
	}

	return true
}
func privateKeyToString(privKey *rsa.PrivateKey) (string) {
	privKeyBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		return ""
	}

	privKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privKeyBytes})
	return string(privKeyPEM)
}

func main(){
	privateKey, publicKey := generateKeyPair()

	// Sign the message with private key
	info, signatureFromSender := signMessage(privateKey)

	result := verifySignature(info, signatureFromSender, publicKey)
	fmt.Println("Verification result : ", result)
}