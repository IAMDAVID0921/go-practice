package main

import (
	"context"

	pb "github.com/AlexsJones/toolbox/grpc-demo/proto"
)

// unary server function
// rpc SayHello (NoParam) returns (HelloResponse);
func (s *helloserver) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
