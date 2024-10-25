package main

import (
    "context"
    "fmt"
    "net"
    "log"

    "github.com/[your/project]/greeter"
    "google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (s *Server) SayGoodbye(ctx context.Context, req *greeter.GoodbyeRequest) (*greeter.GoodbyeReply, error) {
	return &greeter.GoodbyeReply{
		Message: fmt.Sprintf("Good-bye, %s!", req.GetName()),
	}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":500５1")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := grpc.NewServer()
	greeter.RegisterGreeterServer(server, &Server{})
	server.Serve(conn)
}