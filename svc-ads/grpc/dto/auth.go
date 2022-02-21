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

func (c *pbAuthConvert) ToGetUserIDByUsernameRequest(username string) *pbAuth.GetUserIDByUsernameRequest {
	return &pbAuth.GetUserIDByUsernameRequest{
		Username: username,
	}
}

func (c *pbAuthConvert) FromGetUserIDByUsernameResponse(in *pbAuth.GetUserIDByUsernameResponse) uint32 {
	return in.Id
}
