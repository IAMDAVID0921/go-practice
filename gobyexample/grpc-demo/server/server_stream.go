package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AlexsJones/toolbox/grpc-demo/proto"
)

func (s *helloserver) SayHelloServerStreaming(ctx context.Context, req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
