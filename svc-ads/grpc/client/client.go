package client

import (
	log "github.com/sirupsen/logrus"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/grpc"
)

type Connection struct {
	Auth *grpc.ClientConn
}

type Client struct {
	AuthService pbAuth.AuthServiceClient
}

func OpenConnection(adsAddress, authAddress string) (*Connection, error) {
	connAuth, err := grpc.Dial(authAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &Connection{
		Auth: connAuth,
	}, nil
}

func (c *Connection) Close() {
	if err := c.Auth.Close(); err != nil {
		log.WithError(err).Error("failed to close Auth connection")
	}
}

func NewClient(conn *Connection) *Client {
	return &Client{
		AuthService: pbAuth.NewAuthServiceClient(conn.Auth),
	}
}
