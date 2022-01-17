package server

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"github.com/u-shylianok/ad-service/svc-auth/service"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	Service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{Service: service}
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.User, error) {
	logrus.WithField("in", in).Error("SignUp")
	return nil, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	logrus.WithField("in", in).Error("SignUp")
	return nil, nil
}

func (s *Server) ParseToken(ctx context.Context, in *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	logrus.WithField("in", in).Error("ParseToken")
	return nil, nil
}

func (s *Server) Identify(ctx context.Context, in *pb.IdentifyRequest) (*pb.IdentifyResponse, error) {
	logrus.WithField("in", in).Error("SignUp")
	return nil, nil
}

func (s *Server) GetUserID(ctx context.Context, in *pb.GetUserIDRequest) (*pb.GetUserIDResponse, error) {
	logrus.WithField("in", in).Error("SignUp")
	return nil, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	logrus.WithField("in", in).Error("SignUp")
	return nil, nil
}
