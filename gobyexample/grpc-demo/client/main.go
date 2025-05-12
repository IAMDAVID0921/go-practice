package main

import (
	"log"

	pb "github.com/AlexsJones/toolbox/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// suppose to make request to server (make insecure request to server, not use any secure credentials)
// which means we need to interact on the same port

const (
	port = ":8080"
)

func main() {
	// same as grpc.Dial (deprecated)
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close() // after making request, at the end we want to close the connection

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
