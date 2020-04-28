package main

import (
	"flag"
	"log"
)

// A very simple go program implementing the hello protobuf.
// It can act as both a server and a client.
// Which kind is determined by the flag, by default it acts as a server.
// To act as a client a "-server=<url>" flag must be passed.

func main() {
	serverURL := flag.String("server", "", "Set the URL to the server if it should act as a client")
	flag.Parse()
	if serverURL == nil || len(*serverURL) == 0 {
		StartServer()
	} else {
		StartClient(*serverURL)
	}
	log.Printf("Done!")
}
