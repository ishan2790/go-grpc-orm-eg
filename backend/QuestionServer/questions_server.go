package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ishan2790/go-grpc-orm-eg/backend/questions"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type QuestionServer struct {
	pb.UnimplementedQuestionServiceServer
}

func (s *QuestionServer) Add(ctx context.Context, in *pb.AddQuestionRequest) (*pb.AddQuestionResponse, error) {
	log.Printf("Received :%v", in.GetName())
	return &pb.AddQuestionResponse{Id: 1}, nil
}

func (s *QuestionServer) Get(ctx context.Context, in *pb.GetQuestionRequest) (*pb.GetQuestionResponse, error) {
	log.Printf("Received get id:%v", in.GetId())
	return &pb.GetQuestionResponse{Id: 1, Name: "Test",Category:"Math", Subcategory:"Average",Difficulty:0, Type:0, Options: []*pb.Opt{{Id:1 ,Name:"Answer1",Flag:true }}}, nil
}

func (s *QuestionServer) GetAll(ctx context.Context, in *pb.Empty) (*pb.GetAllQuestionResponse, error) {
	log.Printf("Received nothing for getall")
	// questionStruct1 := &pb.GetQuestionResponse{Id: 1, Name: "Delhi",Category:"Math", Subcategory:"Average",Difficulty:0, Type:0, Options: []*pb.Opt{{Id:1 ,Name:"Answer",Flag:true }}};
	// questionStruct2 := &pb.GetQuestionResponse{Id: 1, Name: "Australia",Category:"Math", Subcategory:"Average",Difficulty:0, Type:0,Options: []*pb.Opt{{Id:1 ,Name:"Answer1",Flag:true }}};

	// array := make([]pb.GetAllQuestionResponse,2)
	// array = append(array, questionStruct1)
	// array = append(array, questionStruct2)
	return &pb.GetAllQuestionResponse{
		Questionlist : []*pb.GetQuestionResponse{
			{Id: 1, Name: "Delhi",Category:"Math", Subcategory:"Average",Difficulty:0, Type:0, Options: []*pb.Opt{{Id:1 ,Name:"Answer",Flag:true }}},
			{Id: 1, Name: "England",Category:"Math", Subcategory:"Average",Difficulty:0, Type:0, Options: []*pb.Opt{{Id:1 ,Name:"London",Flag:true }}},
		},
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("FAiled to start server: %v",err)
	}

	serv := grpc.NewServer()

	pb.RegisterQuestionServiceServer(serv, &QuestionServer{})

	if err:= serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v",err)
	}
}