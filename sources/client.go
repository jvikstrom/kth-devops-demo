package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/jvikstrom/devops-demo/sources/proto"
	"google.golang.org/grpc"
)

func StartClient(url string) {
	log.Printf("Starting client towards server at: %v", url)
	if err := runClient(url); err != nil {
		log.Fatalf("There was an error running the client: %v", err)
	}
}

// Maximum number of requests a client can have that have not been responded to.
const maxOutstanding = 20

// Maximum number of requests we permit fail before we terminate the client.
const maxFails = 10

func runClient(url string) error {
	// Connect to our server.
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := proto.NewHelloServiceClient(conn)

	nfailed := 0
	failChan := make(chan bool)
	doneChan := make(chan struct{}, maxOutstanding)
	var shouldQuit int32
	go func() {
		// Handles failed requests.
		for {
			<-failChan
			nfailed++
			if nfailed > maxFails {
				atomic.StoreInt32(&shouldQuit, 1)
			}
		}
	}()

	startTime := time.Now()
	// Send first batch of requests.
	for i := 0; i < maxOutstanding; i++ {
		sendRequestAsync(client, doneChan, failChan)
	}

	// Resend requests forever.
	for {
		<-doneChan // Wait for a request to notify us it's done.
		if atomic.LoadInt32(&shouldQuit) == 1 {
			totalTimeMs := time.Now().Sub(startTime).Milliseconds()
			return fmt.Errorf("Too many requests have failed, ran for %v ms, exiting", totalTimeMs)
		}
		sendRequestAsync(client, doneChan, failChan) // Send new request to make sure we saturate the number of outstanding requests.
	}

}

// Helper to send requests in async to make the code nicer.
func sendRequestAsync(client proto.HelloServiceClient, doneChan chan struct{}, failChan chan bool) {
	go func() {
		if err := sendRequest(client); err != nil {
			log.Printf("Error sending request: %v", err)
			failChan <- true
		}
		doneChan <- struct{}{} // Mark request as done.
	}()
}

func sendRequest(client proto.HelloServiceClient) error {
	n := rand.Int63() % 1000                                                               // Get random number in range [0,1000).
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second)) // Have a 5 second timeout for each request
	req := &proto.SayHelloRequest{Start: n}
	_, err := client.SayHello(ctx, req)
	cancel()
	return err
}
