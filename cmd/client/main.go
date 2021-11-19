package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yusufpapurcu/twirp-test/rpc/haberdasher"
)

func main() {
	client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 12})
	if err != nil {
		panic(err)
	}
	fmt.Printf("I have a nice new hat: %+v", hat)
}
