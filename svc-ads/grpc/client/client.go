package client

import (
	log "github.com/sirupsen/logrus"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	authConn    *grpc.ClientConn
	AuthService pbAuth.AuthServiceClient
}

func New(authAddress string) (*Client, error) {
	newClient := &Client{}

	authConn, err := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	newClient.AuthService = pbAuth.NewAuthServiceClient(authConn)
	newClient.authConn = authConn

	return newClient, nil
}

func (c *Client) Close() {
	if err := c.authConn.Close(); err != nil {
		log.WithError(err).Error("failed to close Auth connection")
	}
}
