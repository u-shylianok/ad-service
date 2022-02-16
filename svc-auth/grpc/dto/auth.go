package dto

import (
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"github.com/u-shylianok/ad-service/svc-auth/domain/model"
)

type pbAuthConvert struct{}

var PbAuth pbAuthConvert

func (c *pbAuthConvert) FromSignUpRequest(req *pbAuth.SignUpRequest) model.User {
	return model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}

func (c *pbAuthConvert) ToUserResponse(user model.UserResponse) *pbAuth.UserResponse {
	return &pbAuth.UserResponse{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

func (c *pbAuthConvert) FromSignInRequest(req *pbAuth.SignInRequest) (string, string) {
	return req.Username, req.Password
}

func (c *pbAuthConvert) ToSignInResponse(token string, expiresAt int64) *pbAuth.SignInResponse {
	return &pbAuth.SignInResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}
}

func (c *pbAuthConvert) FromParseTokenRequest(req *pbAuth.ParseTokenRequest) string {
	return req.Token
}

func (c *pbAuthConvert) ToParseTokenResponse(userID uint32) *pbAuth.ParseTokenResponse {
	return &pbAuth.ParseTokenResponse{
		UserId: userID,
	}
}

func (c *pbAuthConvert) FromGetUserRequest(req *pbAuth.GetUserRequest) uint32 {
	return req.Id
}

func (c *pbAuthConvert) ToGetUserResponse(user model.UserResponse) *pbAuth.GetUserResponse {
	return &pbAuth.GetUserResponse{
		User: c.ToUserResponse(user),
	}
}

func (c *pbAuthConvert) FromListUsersInIDsRequest(req *pbAuth.ListUsersInIDsRequest) []uint32 {
	return req.Ids
}

func (c *pbAuthConvert) ToListUsersInIDsResponse(users []model.UserResponse) *pbAuth.ListUsersInIDsResponse {
	result := make([]*pbAuth.UserResponse, len(users))
	for i, user := range users {
		result[i] = c.ToUserResponse(user)
	}
	return &pbAuth.ListUsersInIDsResponse{
		Users: result,
	}
}
