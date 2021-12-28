package main

import (
	"log"

	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ad"
)

func main() {

	adsClient := pbAds.NewAdServiceClient(nil)

	log.Println(adsClient)
}
