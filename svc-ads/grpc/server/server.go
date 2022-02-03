package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	"github.com/u-shylianok/ad-service/svc-ads/grpc/client"
	"github.com/u-shylianok/ad-service/svc-ads/grpc/dto"
	"github.com/u-shylianok/ad-service/svc-ads/service"
)

type Server struct {
	pb.UnimplementedAdServiceServer

	Service *service.Service
	clients *client.Client
}

func New(service *service.Service, clients *client.Client) *Server {
	return &Server{
		Service: service,
		clients: clients,
	}
}

func (s *Server) GetAd(ctx context.Context, in *pb.GetAdRequest) (*pb.GetAdResponse, error) {
	ad, err := s.Service.GetAd(dto.FromPbAds_GetAdRequest(in))
	if err != nil {
		return nil, err
	}

	user, err := s.clients.AuthService.GetUser(context.Background(), dto.ToPbAuth_GetUserRequest(ad.UserID))
	if err != nil {
		return nil, err
	}
	return dto.ToPbAds_GetAdResponse(ad, user), nil
}

func (s *Server) ListAds(ctx context.Context, in *pb.ListAdsRequest) (*pb.ListAdsResponse, error) {
	ads, err := s.Service.ListAds(dto.FromPbAds_ListAdsRequest(in))
	if err != nil {
		return nil, err
	}

	ids := make(map[uint32]bool)
	usersIDs := []uint32{}
	for _, ad := range ads {
		usersID := ad.UserID
		if _, value := ids[usersID]; !value {
			ids[usersID] = true
			usersIDs = append(usersIDs, usersID)
		}
	}
	users, err := s.clients.AuthService.ListUsersInIDs(context.Background(), dto.ToPbAuth_ListUsersInIDsRequest(usersIDs))
	if err != nil {
		return nil, err
	}
	return dto.ToPbAds_ListAdsResponse(ads, users), nil
}

func (s *Server) SearchAds(ctx context.Context, in *pb.SearchAdsRequest) (*pb.SearchAdsResponse, error) {
	ads, err := s.Service.SearchAds(dto.FromPbAds_SearchAdsRequest(in))
	if err != nil {
		return nil, err
	}

	ids := make(map[uint32]bool)
	usersIDs := []uint32{}
	for _, ad := range ads {
		usersID := ad.UserID
		if _, value := ids[usersID]; !value {
			ids[usersID] = true
			usersIDs = append(usersIDs, usersID)
		}
	}
	users, err := s.clients.AuthService.ListUsersInIDs(context.Background(), dto.ToPbAuth_ListUsersInIDsRequest(usersIDs))
	if err != nil {
		return nil, err
	}
	return dto.ToPbAds_SearchAdsResponse(ads, users), nil
}

func (s *Server) CreateAd(ctx context.Context, in *pb.CreateAdRequest) (*pb.Ad, error) {
	adID, err := s.Service.CreateAd(dto.FromPbAds_CreateAdRequest(in))
	if err != nil {
		return nil, err
	}

	// TODO : draft part
	ad, err := s.Service.GetAd(adID, model.AdsOptional{})
	if err != nil {
		return nil, err
	}
	//
	return dto.ToPbAds_Ad(ad), nil
}

func (s *Server) UpdateAd(ctx context.Context, in *pb.UpdateAdRequest) (*pb.Ad, error) {
	adID, err := s.Service.UpdateAd(dto.FromPbAds_UpdateAdRequest(in))
	if err != nil {
		return nil, err
	}

	// TODO : draft part
	ad, err := s.Service.GetAd(adID, model.AdsOptional{})
	if err != nil {
		return nil, err
	}
	//
	return dto.ToPbAds_Ad(ad), nil
}

func (s *Server) DeleteAd(ctx context.Context, in *pb.DeleteAdRequest) (*empty.Empty, error) {
	if err := s.Service.DeleteAd(dto.FromPbAds_DeleteAdRequest(in)); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) ListPhotos(ctx context.Context, in *pb.ListPhotosRequest) (*pb.ListPhotosResponse, error) {
	adID := dto.FromPbAds_ListPhotosRequest(in)

	var result []string
	if adID == 0 {
		photos, err := s.Service.ListPhotos()
		if err != nil {
			return nil, err
		}
		result = photos
	} else {
		photos, err := s.Service.ListAdPhotos(adID)
		if err != nil {
			return nil, err
		}
		result = photos
	}
	return dto.ToPbAds_ListPhotosResponse(result), nil
}

func (s *Server) ListTags(ctx context.Context, in *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	adID := dto.FromPbAds_ListTagsRequest(in)

	var result []string
	if adID == 0 {
		tags, err := s.Service.ListTags()
		if err != nil {
			return nil, err
		}
		result = tags
	} else {
		tags, err := s.Service.ListAdTags(adID)
		if err != nil {
			return nil, err
		}
		result = tags
	}
	return dto.ToPbAds_ListTagsResponse(result), nil
}
