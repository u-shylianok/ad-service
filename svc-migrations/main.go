package main

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

func main() {

	setupGlobalLogger()
	log.Info("start migrations")
	{
		log.Info("start migrations AUTH")
		dbAuth, err := NewPostgresDB(Config{
			Host:     os.Getenv("DB_AUTH_HOST"),
			Port:     os.Getenv("DB_AUTH_PORT"),
			Username: os.Getenv("DB_AUTH_USER"),
			Password: os.Getenv("DB_AUTH_PASSWORD"),
			DBName:   os.Getenv("DB_AUTH_DBNAME"),
			SSLMode:  os.Getenv("DB_AUTH_SSL"),
		})
		if err != nil {
			log.Fatalf("failed to initialize auth db: %s", err)
		}

		log.Info("try to do migrations")
		if err := goose.Up(dbAuth.DB, "migrations/db-auth"); err != nil {
			log.Fatal("Error occurred in migration:", err)
		}
		log.Info("end migrations AUTH")
	}

	// {
	// 	log.Info("start migrations ADS")
	// 	dbAds, err := NewPostgresDB(Config{
	// 		Host:     os.Getenv("DB_ADS_HOST"),
	// 		Port:     os.Getenv("DB_ADS_PORT"),
	// 		Username: os.Getenv("DB_ADS_USER"),
	// 		Password: os.Getenv("DB_ADS_PASSWORD"),
	// 		DBName:   os.Getenv("DB_ADS_DBNAME"),
	// 		SSLMode:  os.Getenv("DB_ADS_SSL"),
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("failed to initialize ads db: %s", err)
	// 	}

	// 	log.Info("try to do migrations")
	// 	if err := goose.Up(dbAds.DB, "migrations/db-ads"); err != nil {
	// 		log.Fatal("Error occurred in migration:", err)
	// 	}
	// 	log.Info("end migrations ADS")
	// }
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

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
