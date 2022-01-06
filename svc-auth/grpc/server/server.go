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

func (s *Server) SignUp(context.Context, *pb.SignUpRequest) (*pb.User, error) {
	logrus.Fatal("SignUp")
	return nil, nil
}

func (s *Server) SignIn(context.Context, *pb.SignInRequest) (*pb.SignInResponse, error) {
	logrus.Fatal("SignIn")
	return nil, nil
}

func (s *Server) Identify(context.Context, *pb.IdentifyRequest) (*pb.IdentifyResponse, error) {
	logrus.Fatal("Identify")
	return nil, nil
}

func (s *Server) GetUserID(context.Context, *pb.GetUserIDRequest) (*pb.GetUserIDResponse, error) {
	logrus.Fatal("GetUserID")
	return nil, nil
}

func (s *Server) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	logrus.Fatal("GetUser")
	return nil, nil
}
