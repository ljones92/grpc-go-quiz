package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/ljones92/grpc-go-quiz"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetQuestion(ctx context.Context, in *pb.QuestionRequest) (*pb.QuestionResponse, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	question := Questions[rand.Intn(len(Questions))]

	return &pb.QuestionResponse{Question: question.Question, Answer: question.Answer}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterQuizServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
