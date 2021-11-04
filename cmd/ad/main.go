package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/handler"
	"github.com/u-shylianok/ad-service/internal/repository"
	"github.com/u-shylianok/ad-service/internal/service"
)

const (
	MAX_HEADER_BYTES = 1 << 20 // 1 MB
	READ_TIMEOUT     = 10 * time.Second
	WRITE_TIMEOUT    = 10 * time.Second
	DEBUG_LOG_LEVEL  = "debug"
)

type Server struct {
	httpServer *http.Server
}

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
	handlers := handler.NewHandler(services)

	srv := new(Server)
	go func() {
		if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err)
		}
	}()

	log.Info("ads service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("ads service shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err)
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err)
	}
	return
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: MAX_HEADER_BYTES,
		ReadTimeout:    READ_TIMEOUT,
		WriteTimeout:   WRITE_TIMEOUT,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func setupGlobalLogger() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == DEBUG_LOG_LEVEL {
		log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.Infof("logger started with log_level = %s", logLevel)
}
