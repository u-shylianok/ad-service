package dto

import pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"

func ToPbAuth_GetUserRequest(userID uint32) *pbAuth.GetUserRequest {
	return &pbAuth.GetUserRequest{
		Id: userID,
	}
}

func ToPbAuth_ListUsersInIDsRequest(usersIDs []uint32) *pbAuth.ListUsersInIDsRequest {
	return &pbAuth.ListUsersInIDsRequest{
		Ids: usersIDs,
	}
}
