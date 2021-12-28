package grpc

import (
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/grpc"
)

type Clients struct {
	AuthService *pbAuth.AuthServiceClient
}

func NewClients(authConn *grpc.ClientConn) *Clients {

	return &Clients{
		AuthService: NewAuthClient(authConn),
	}
}

func NewAuthClient(conn *grpc.ClientConn) *pbAuth.AuthServiceClient {
	authClient := pbAuth.NewAuthServiceClient(conn)
	return &authClient
}
