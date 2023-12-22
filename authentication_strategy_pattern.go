package main

import (
	"context"
	"fmt"
)

type AuthenticatorStrategy interface {
	Authenticate(ctx context.Context, token []byte) error
}

type Authn struct {
	strategy AuthenticatorStrategy
}

func main (){
	fmt.Println("Hello from the world !")
}

type JWTStrategy struct {}
type ApiStrategy struct {}
type ShareKeyStrategy struct{}

func (svc *JWTStrategy) Authenticate(ctx context.Context, token []byte) (error) {
	return nil
}

func (svc *ApiStrategy) Authenticate(ctx context.Context, token []byte) (error) {
	return nil
}

func (svc *ShareKeyStrategy) Authenticate(ctx context.Context, token []byte) (error) {
	return nil
}
