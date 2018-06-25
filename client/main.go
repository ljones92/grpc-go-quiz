package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
		answer, _ := reader.ReadString('\n')

		trimmedAnswer := strings.TrimRight(answer, "\n")

		if trimmedAnswer == q.Answer {
			fmt.Println("Correct!!!")
		} else {
			fmt.Println("That's wrong the correct answer is: " + q.Answer)
		}
	}
}
