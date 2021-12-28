package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	pb "github.com/u-shylianok/ad-service/svc-ads/client/ad"
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

	s := grpc.NewServer()
	pb.RegisterAdServiceServer(s, nil)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	s.GracefulStop()
}
