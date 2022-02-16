package dto

import pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"

type pbAuthConvert struct{}

var PbAuth pbAuthConvert

func (c *pbAuthConvert) ToGetUserRequest(userID uint32) *pbAuth.GetUserRequest {
	return &pbAuth.GetUserRequest{
		Id: userID,
	}
}

func (c *pbAuthConvert) ToListUsersInIDsRequest(usersIDs []uint32) *pbAuth.ListUsersInIDsRequest {
	return &pbAuth.ListUsersInIDsRequest{
		Ids: usersIDs,
	}
}
