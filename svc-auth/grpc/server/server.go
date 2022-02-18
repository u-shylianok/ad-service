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
	userID, err := s.Service.CreateUser(dto.PbAuth.FromSignUpRequest(in))
	if err != nil {
		return nil, err
	}

	// TODO : draft part
	user, err := s.Service.GetUser(userID)
	if err != nil {
		return nil, err
	}
	//
	return dto.PbAuth.ToUserResponse(user), nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	id, err := s.Service.CheckUser(dto.PbAuth.FromSignInRequest(in))
	if err != nil {
		return nil, err
	}

	token, expiresAt, err := s.Service.GenerateToken(id)
	if err != nil {
		return nil, err
	}
	return dto.PbAuth.ToSignInResponse(token, expiresAt), nil
}

func (s *Server) ParseToken(ctx context.Context, in *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	userID, err := s.Service.ParseToken(dto.PbAuth.FromParseTokenRequest(in))
	if err != nil {
		return nil, err
	}
	return dto.PbAuth.ToParseTokenResponse(userID), nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.Service.GetUser(dto.PbAuth.FromGetUserRequest(in))
	if err != nil {
		return nil, err
	}
	return dto.PbAuth.ToGetUserResponse(user), nil
}

func (s *Server) GetUserIDByUsername(ctx context.Context, in *pb.GetUserIDByUsernameRequest) (*pb.GetUserIDByUsernameResponse, error) {
	users, err := s.Service.GetUserIDByUsername(dto.PbAuth.FromGetUserIDByUsernameRequest(in))
	if err != nil {
		return nil, err
	}
	return dto.PbAuth.ToGetUserIDByUsernameResponse(users), nil
}

func (s *Server) ListUsersInIDs(ctx context.Context, in *pb.ListUsersInIDsRequest) (*pb.ListUsersInIDsResponse, error) {
	users, err := s.Service.ListUsersInIDs(dto.PbAuth.FromListUsersInIDsRequest(in))
	if err != nil {
		return nil, err
	}
	return dto.PbAuth.ToListUsersInIDsResponse(users), nil
}
