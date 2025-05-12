package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/AlexsJones/toolbox/grpc-demo/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	// create a channel
	// one routine

	waitc := make(chan struct{})

	// your goroutine
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming, %v", err)
			}
			log.Println(message)
		}
		// close channel
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the request: %v", err)
		}
		time.Sleep(time.Second * 2)
	}
	stream.CloseSend()
	<-waitc

	log.Printf("Bidirectional streaming finished.")
}
