package dto

import (
	"github.com/u-shylianok/ad-service/api-gateway/domain/model"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
)

type pbAuthConvert struct{}

var PbAuth pbAuthConvert

func (c *pbAuthConvert) ToSignUpRequest(req model.SignUpRequest) *pbAuth.SignUpRequest {
	return &pbAuth.SignUpRequest{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}

func (c *pbAuthConvert) ToSignInRequest(req model.SignInRequest) *pbAuth.SignInRequest {
	return &pbAuth.SignInRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func (c *pbAuthConvert) FromUser(user *pbAuth.UserResponse) model.UserResponse {
	return model.UserResponse{
		Name:     user.GetName(),
		Username: user.GetUsername(),
	}
}
