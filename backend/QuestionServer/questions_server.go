package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ishan2790/go-grpc-orm-eg"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type QuestionServer struct {
	pb.UnimplementedQuestionServer
}

func (s *QuestionServer) Add(ctx context.Context, in *pb.AddQuestionRequest) (*pb.AddQuestionResponse, error) {
	log.Printf("Received :%v", in.GetName())
	return &pb.AddQuestionResponse{Id: 1}, nil
}

func (s *QuestionServer) Get(ctx context.Context, in *pb.GetQuestionRequest) (*pb.GetQuestionResponse, error) {
	log.Printf("Received get id:%v", in.GetId())
	return &pb.GetQuestionResponse{Id: 1, Name: "Test",Category:"Math", SubCategory:"Average",Difficulty:"Easy", Type:0, Options: {Id:1 ,Name:"Answer1",flag:true }}, nil
}

func (s *QuestionServer) GetAll(ctx context.Context) (*[]pb.GetQuestionResponse, error) {
	log.Printf("Received nothing for getall")
	array := {{Id: 1, Name: "Test",Category:"Math", SubCategory:"Average",Difficulty:"Easy", Type:0, Options: {Id:1 ,Name:"Answer1",flag:true }}, 
	{Id: 2, Name: "Test1",Category:"Math", SubCategory:"Average",Difficulty:"Easy", Type:0, Options: {Id:1 ,Name:"Answer2",flag:true }}}
	return &array, nil
}


func main() {
	lis, err : = net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("FAiled to start server: %v",err)
	}

	s := grpc.NewServer()

	pb.RegisterQuestionServer(s, &QuestionServer{})

	if err:= s.Serve(lis); err != nik {
		log.Fatalf("Failed to serve %v",err)
	}
}