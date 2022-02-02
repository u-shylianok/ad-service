package dto

import (
	"github.com/u-shylianok/ad-service/api-gateway/domain/model"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
)

func ToPbAuth_SignUpRequest(req model.SignUpRequest) *pbAuth.SignUpRequest {
	return &pbAuth.SignUpRequest{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}

func ToPbAuth_SignInRequest(req model.SignInRequest) *pbAuth.SignInRequest {
	return &pbAuth.SignInRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func FromPbAuth_User(user *pbAuth.UserResponse) model.UserResponse {
	return model.UserResponse{
		Name:     user.Name,
		Username: user.Username,
	}
}
