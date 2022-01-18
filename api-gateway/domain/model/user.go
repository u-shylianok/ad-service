package model

import pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (in *SignUpRequest) ToPb() *pbAuth.SignUpRequest {
	return &pbAuth.SignUpRequest{
		Name:     in.Name,
		Username: in.Username,
		Password: in.Password,
	}
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (in *SignInRequest) ToPb() *pbAuth.SignInRequest {
	return &pbAuth.SignInRequest{
		Username: in.Username,
		Password: in.Password,
	}
}
