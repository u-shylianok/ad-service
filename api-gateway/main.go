package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/api-gateway/grpc/client"
	"github.com/u-shylianok/ad-service/api-gateway/handler"
)

const (
	MaxHeaderBytes = 1 << 20 // 1 MB
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func main() {
	setupGlobalLogger()

	log.Info("start gRPC clients connection")
	grpcClients, err := client.New(os.Getenv("SVC_ADS_ADDRESS"), os.Getenv("SVC_AUTH_ADDRESS"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcClients.Close()
	log.Info("connection started")

	handlers := handler.NewHandler(grpcClients)

	srv := new(Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err)
		}
	}()

	log.Info("api-gateway started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("api-gateway shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err)
	}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: MaxHeaderBytes,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
	}
	log.Println(s)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
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
