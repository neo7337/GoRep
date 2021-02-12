package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/neo7337/goRep/my-grpc/stub"
)

func main() {

	fmt.Printf("Running gRPC Server\n")

	serv, err := net.Listen("tcp", ":4010")
	if err != nil {
		log.Fatalf("Error on Starting Server %v", err)
	}

	s := stub.Server{}

	grpcServer := grpc.NewServer()

	stub.RegisterHelloServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(serv); err != nil {
		log.Fatalf("gRPC Server did not start: %s", err)
	}
}
