package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/ljones92/grpc-go-quiz"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewQuizClient(conn)

	for {
		q, err := c.GetQuestion(context.Background(), &pb.QuestionRequest{})
		if err != nil {
			log.Fatalf("could not get question: %v", err)
		}

		fmt.Println(q.Question)
		fmt.Println("Enter your answer...")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		s, err := c.CheckAnswer(context.Background(), &pb.Answer{
			Body:       text,
			QuestionId: q.Id,
		})
		if err != nil {
			log.Fatalf("could not check answer: %v", err)
		}

		if s.Correct {
			fmt.Println("Correct!!!")
		} else {
			fmt.Println("That's wrong the correct answer is: " + s.Answer)
		}
	}
}
