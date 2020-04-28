package main

import (
	"context"
	"log"
	"net"

	proto "github.com/jvikstrom/devops-demo/sources/proto"
	"google.golang.org/grpc"
)

func StartServer() {
	hello := helloServer{}
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, hello)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error running GRPC server: %v", err)
	}
}

type helloServer struct {
}

// SayHello is a handler for the sayHello rpc in the Hello service.
// Just calculates req.Start primes in an inefficient way to use CPU.
func (h helloServer) SayHello(ctx context.Context, req *proto.SayHelloRequest) (res *proto.SayHelloResponse, err error) {
	// Let's just calculate req.Start number of primes in N^2.
	primeList := []int{}
	n := 2
	for len(primeList) < int(req.Start) {
		divisable := false
		for i := 2; i < n; i++ {
			if n%i == 0 {
				divisable = true
				break
			}
		}
		if !divisable {
			primeList = append(primeList, n)
		}
	}
	res = &proto.SayHelloResponse{End: int64(len(primeList))}
	return
}
