package client

import (
	log "github.com/sirupsen/logrus"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	adsConn     *grpc.ClientConn
	authConn    *grpc.ClientConn
	AdsService  pbAds.AdServiceClient
	AuthService pbAuth.AuthServiceClient
}

func New(adsAddress, authAddress string) (*Client, error) {
	newClient := &Client{}

	adsConn, err := grpc.Dial(adsAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	authConn, err := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	newClient.AdsService = pbAds.NewAdServiceClient(adsConn)
	newClient.AuthService = pbAuth.NewAuthServiceClient(authConn)
	newClient.adsConn = adsConn
	newClient.authConn = authConn

	return newClient, nil
}

func (c *Client) Close() {
	if err := c.adsConn.Close(); err != nil {
		log.WithError(err).Error("failed to close Ad connection")
	}
	if err := c.authConn.Close(); err != nil {
		log.WithError(err).Error("failed to close Auth connection")
	}
}
