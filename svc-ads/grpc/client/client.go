package client

import (
	log "github.com/sirupsen/logrus"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/grpc"
)

type Client struct {
	authConn    *grpc.ClientConn
	AuthService pbAuth.AuthServiceClient
}

func New(authAddress string) (*Client, error) {
	var newClient *Client

	log.Info("start dial auth " + authAddress)
	authConn, err := grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	log.WithField("auth", authConn).Info("auth connection")

	newClient.AuthService = pbAuth.NewAuthServiceClient(authConn)
	newClient.authConn = authConn

	return newClient, nil
}

func (c *Client) Close() {
	if err := c.authConn.Close(); err != nil {
		log.WithError(err).Error("failed to close Auth connection")
	}
}
