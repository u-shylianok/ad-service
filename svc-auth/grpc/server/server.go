package server

import (
	"context"

	pb "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"github.com/u-shylianok/ad-service/svc-auth/grpc/dto"
	"github.com/u-shylianok/ad-service/svc-auth/service"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	Service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{Service: service}
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.UserResponse, error) {
	userID, err := s.Service.CreateUser(dto.FromPbAuth_SignUpRequest(in))
	if err != nil {
		return nil, err
	}

	// TODO : draft part
	user, err := s.Service.GetUser(userID)
	if err != nil {
		return nil, err
	}
	//
	return dto.ToPbAuth_UserResponse(user), nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	id, err := s.Service.CheckUser(dto.FromPbAuth_SignInRequest(in))
	if err != nil {
		return nil, err
	}

	token, expiresAt, err := s.Service.GenerateToken(id)
	if err != nil {
		return nil, err
	}
	return dto.ToPbAuth_SignInResponse(token, expiresAt), nil
}

func (s *Server) ParseToken(ctx context.Context, in *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	userID, err := s.Service.ParseToken(dto.FromPbAuth_ParseTokenRequest(in))
	if err != nil {
		return nil, err
	}
	return dto.ToPbAuth_ParseTokenResponse(userID), nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.Service.GetUser(dto.FromPbAuth_GetUserRequest(in))
	if err != nil {
		return nil, err
	}
	return dto.ToPbAuth_GetUserResponse(user), nil
}
