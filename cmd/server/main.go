package main

import (
	"net/http"

	"github.com/yusufpapurcu/twirp-test/internal/haberdasherserver"
	"github.com/yusufpapurcu/twirp-test/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
