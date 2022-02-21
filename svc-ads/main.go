package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
	pb "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/grpc/client"
	"github.com/u-shylianok/ad-service/svc-ads/grpc/server"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
	"github.com/u-shylianok/ad-service/svc-ads/service"
	"google.golang.org/grpc"
)

func main() {
	setupGlobalLogger()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	log.Info("start gRPC clients connection")
	grpcClients, err := client.New(os.Getenv("SVC_AUTH_ADDRESS"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcClients.Close()
	log.Info("gRPC connection started")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	srv := grpc.NewServer()
	go func() {
		lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		server := server.New(services, grpcClients)
		pb.RegisterAdServiceServer(srv, server)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Info("svc-ads started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	srv.GracefulStop()
	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err)
	}
}

func setupGlobalLogger() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" && logLevel != "default" {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			log.WithError(err).Error("failed to parse log level from env")
		} else {
			log.SetLevel(level)
		}
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.WithField("log_level", logLevel).Info("logger initialised")
}
