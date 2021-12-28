package grpc

import (
	"context"

	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
)

type Server struct{}

func (s *Server) SignUp(context.Context, *pbAuth.SignUpRequest) (*pbAuth.User, error) {
	return nil, nil
}

func (s *Server) SignIn(context.Context, *pbAuth.SignInRequest) (*pbAuth.SignInResponse, error) {
	return nil, nil
}

func (s *Server) Identify(context.Context, *pbAuth.IdentifyRequest) (*pbAuth.IdentifyResponse, error) {
	return nil, nil
}

func (s *Server) GetUserID(context.Context, *pbAuth.GetUserIDRequest) (*pbAuth.GetUserIDResponse, error) {
	return nil, nil
}
