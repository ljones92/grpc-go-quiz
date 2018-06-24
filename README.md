# gRPC Quiz

This is a very simple quiz command line tool written in Go using gRPC, the server will grap random questions from the server/questions.go file and wait for an answer from the client

## Running

Execute the following commands to run this locally:

`go run server/main.go server/questions.go`

and then in a separate terminal:

`go run client/main.go`
