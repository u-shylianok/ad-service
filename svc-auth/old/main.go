package main

import (
	"log"

	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
)

func main() {
	adsClient := pbAuth.NewAuthServiceClient(nil)

	log.Println(adsClient)
}
