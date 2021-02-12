package main

import (
	"context"
	"log"

	"github.com/neo7337/goRep/my-grpc/stub"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No Connection Established: %s", err)
	}

	defer conn.Close()

	c := stub.NewHelloServiceClient(conn)
	response, err := c.Consume(context.Background(), &stub.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error while calling Stub Consume: %s", err)
	}
	log.Printf("Response from Server: %s", response.Body)

}
