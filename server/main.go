package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	pb "github.com/ljones92/grpc-go-quiz"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetQuestion(ctx context.Context, in *pb.QuestionRequest) (*pb.Question, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	question := Questions[rand.Intn(len(Questions))]
	return &pb.Question{Question: question.Body, Id: question.ID}, nil
}

func (s *server) CheckAnswer(ctx context.Context, in *pb.Answer) (*pb.AnswerCorrect, error) {
	givenAnswer := strings.TrimRight(in.Body, "\n")
	correctAnswer := func() string {
		for i := range Questions {
			if Questions[i].ID == in.QuestionId {
				return Questions[i].Answer
			}
		}

		return ""
	}()

	return &pb.AnswerCorrect{Correct: givenAnswer == correctAnswer, Answer: correctAnswer}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQuizServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
