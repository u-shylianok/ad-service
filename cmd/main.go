package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/u-shylianok/ad-service/internal/server"
)

func main() {
	server.Run()
}
