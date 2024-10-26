package service

import (
	"context"
	"fmt"
	"log"
	"net"

	mongo "root/mongo"
	greeter "root/protobuf/greeter"

	"google.golang.org/grpc"
)

type Server struct {
	greeter.UnimplementedGreeterServer
	db *mongo.DB // Поле db для подключения к базе данных
}

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

func (s *Server) CreateMember(ctx context.Context, req *greeter.CreateMemberRequest) (*greeter.MemberReply, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	member := &greeter.CreateMemberRequest{Name: req.GetName()}
	if err := s.db.InsertMember(member); err != nil {
		return nil, err
	}

	return &greeter.MemberReply{Name: req.GetName()}, nil
}

func RunGRPC() {
	log.Println("start gRPC server at :50051")
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Инициализируем подключение к базе данных
	mongo.Init() // Инициализация подключения

	server := grpc.NewServer()
	greeter.RegisterGreeterServer(server, &Server{db: mongo.Conn}) // Используем mongo.Conn
	server.Serve(conn)
}