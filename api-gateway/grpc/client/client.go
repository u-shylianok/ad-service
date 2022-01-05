package client

import (
	log "github.com/sirupsen/logrus"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/grpc"
)

type Connection struct {
	Ads  *grpc.ClientConn
	Auth *grpc.ClientConn
}

type Client struct {
	Ad   pbAds.AdServiceClient
	Auth pbAuth.AuthServiceClient
}

func OpenConnection(adsAddress, authAddress string) (*Connection, error) {
	connAds, err := grpc.Dial(adsAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	// connAuth, err := grpc.Dial(authAddress, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	return nil, err
	// }
	return &Connection{
		Ads:  connAds,
		Auth: nil,
	}, nil
}

func (c *Connection) Close() {
	if err := c.Ads.Close(); err != nil {
		log.WithError(err).Error("failed to close Ad connection")
	}
	if err := c.Auth.Close(); err != nil {
		log.WithError(err).Error("failed to close Auth connection")
	}
}

func NewClient(conn *Connection) *Client {
	return &Client{
		Ad:   pbAds.NewAdServiceClient(conn.Ads),
		Auth: pbAuth.NewAuthServiceClient(conn.Auth),
	}
}
