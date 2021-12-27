package svc_auth

import pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ad"

func MyFunc() {
	adsClient := pbAds.NewAdServiceClient(nil)

	adsClient.ListAds()
}
