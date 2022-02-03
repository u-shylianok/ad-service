package dto

import (
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"github.com/u-shylianok/ad-service/svc-auth/domain/model"
)

func FromPbAuth_SignUpRequest(req *pbAuth.SignUpRequest) model.User {
	return model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}

func ToPbAuth_UserResponse(user model.UserResponse) *pbAuth.UserResponse {
	return &pbAuth.UserResponse{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

func FromPbAuth_SignInRequest(req *pbAuth.SignInRequest) (string, string) {
	return req.Username, req.Password
}

func ToPbAuth_SignInResponse(token string, expiresAt int64) *pbAuth.SignInResponse {
	return &pbAuth.SignInResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}
}

func FromPbAuth_ParseTokenRequest(req *pbAuth.ParseTokenRequest) string {
	return req.Token
}

func ToPbAuth_ParseTokenResponse(userID uint32) *pbAuth.ParseTokenResponse {
	return &pbAuth.ParseTokenResponse{
		UserId: userID,
	}
}

func FromPbAuth_GetUserRequest(req *pbAuth.GetUserRequest) uint32 {
	return req.Id
}

func ToPbAuth_GetUserResponse(user model.UserResponse) *pbAuth.GetUserResponse {
	return &pbAuth.GetUserResponse{
		User: ToPbAuth_UserResponse(user),
	}
}

func FromPbAuth_ListUsersInIDsRequest(req *pbAuth.ListUsersInIDsRequest) []uint32 {
	return req.Ids
}

func ToPbAuth_ListUsersInIDsResponse(users []model.UserResponse) *pbAuth.ListUsersInIDsResponse {
	result := make([]*pbAuth.UserResponse, len(users))
	for i, user := range users {
		result[i] = ToPbAuth_UserResponse(user)
	}
	return &pbAuth.ListUsersInIDsResponse{
		Users: result,
	}
}
