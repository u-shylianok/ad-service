package server

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/repository"
)

type Server struct {
	httpServer *http.Server
}

func Run() error {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)

	logrus.Infoln(repos.Auth.Get("test", "$2a$10$1hN6TfPRPS9usxbx9DVoY.ix6a8o.kxsednj6CPTkHujR2JGbvLXG"))

	return nil
}

func (s *Server) Run() error {

	return s.httpServer.ListenAndServe()
}
