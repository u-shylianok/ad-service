package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	pb "github.com/u-shylianok/ad-service/svc-ads/client/example"
)

func main() {
	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Username: os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   os.Getenv("DB_DBNAME"),
	// 	SSLMode:  os.Getenv("DB_SSL"),
	// })
	// if err != nil {
	// 	log.Fatalf("failed to initialize db: %s", err)
	// }

	//repos := repository.NewRepository(db)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := NewServer()
	pb.RegisterExampleServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	s.GracefulStop()
}

type Server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *Server) ExampleFunc(ctx context.Context, in *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	value := in.GetValue()
	value *= 2

	return &pb.ExampleResponse{Value: value}, nil
}

func NewServer() *Server {
	return &Server{}
}
