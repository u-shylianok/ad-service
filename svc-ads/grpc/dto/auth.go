package dto

import pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"

func ToPbAuth_GetUserRequest(userID uint32) *pbAuth.GetUserRequest {
	return &pbAuth.GetUserRequest{
		Id: userID,
	}
}
