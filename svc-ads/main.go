package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
	pb "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/grpc/server"
	"github.com/u-shylianok/ad-service/svc-ads/model"
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

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	result, err := services.Ad.GetAd(1, model.AdOptionalFieldsParam{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", result)

	srv := grpc.NewServer()
	go func() {
		lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		server := server.NewServer(services)
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
	if logLevel != "" {
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
