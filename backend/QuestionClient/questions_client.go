package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ishan2790/go-grpc-orm-eg/backend"
	"google.golang.org/grpc"
)


const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure, grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect %v",err)
	}

	defer conn.Close()

	c := pb.GetUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &pb.GetQuestionRequest{Id: 1})

	if err != nil {
		log.Fatalf("Did not get  : %v", err)
	}

	log.Printf(`Question:
	Name:%s`, r.GetName())
}