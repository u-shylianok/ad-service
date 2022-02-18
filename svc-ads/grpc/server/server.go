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
	ad, err := s.Service.GetAd(dto.PbAds.FromGetAdRequest(in))
	if err != nil {
		return nil, err
	}

	user, err := s.clients.AuthService.GetUser(context.Background(), dto.PbAuth.ToGetUserRequest(ad.UserID))
	if err != nil {
		return nil, err
	}
	return dto.PbAds.ToGetAdResponse(ad, user), nil
}

func (s *Server) ListAds(ctx context.Context, in *pb.ListAdsRequest) (*pb.ListAdsResponse, error) {
	ads, err := s.Service.ListAds(dto.PbAds.FromListAdsRequest(in))
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
	users, err := s.clients.AuthService.ListUsersInIDs(context.Background(), dto.PbAuth.ToListUsersInIDsRequest(usersIDs))
	if err != nil {
		return nil, err
	}
	return dto.PbAds.ToListAdsResponse(ads, users), nil
}

func (s *Server) SearchAds(ctx context.Context, in *pb.SearchAdsRequest) (*pb.SearchAdsResponse, error) {
	username, filter := dto.PbAds.FromSearchAdsRequest(in)
	if username != "" {
		pbUserID, err := s.clients.AuthService.GetUserIDByUsername(context.Background(),
			dto.PbAuth.ToGetUserIDByUsernameRequest(username))
		if err != nil {
			return nil, err
		}
		filter.UserID = dto.PbAuth.FromGetUserIDByUsernameResponse(pbUserID)
	}

	ads, err := s.Service.SearchAds(filter)
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
	users, err := s.clients.AuthService.ListUsersInIDs(context.Background(), dto.PbAuth.ToListUsersInIDsRequest(usersIDs))
	if err != nil {
		return nil, err
	}
	return dto.PbAds.ToSearchAdsResponse(ads, users), nil
}

func (s *Server) CreateAd(ctx context.Context, in *pb.CreateAdRequest) (*pb.Ad, error) {
	adID, err := s.Service.CreateAd(dto.PbAds.FromCreateAdRequest(in))
	if err != nil {
		return nil, err
	}

	// TODO : draft part
	ad, err := s.Service.GetAd(adID, model.AdsOptional{})
	if err != nil {
		return nil, err
	}
	//
	return dto.PbAds.ToAd(ad), nil
}

func (s *Server) UpdateAd(ctx context.Context, in *pb.UpdateAdRequest) (*pb.Ad, error) {
	adID, err := s.Service.UpdateAd(dto.PbAds.FromUpdateAdRequest(in))
	if err != nil {
		return nil, err
	}

	// TODO : draft part
	ad, err := s.Service.GetAd(adID, model.AdsOptional{})
	if err != nil {
		return nil, err
	}
	//
	return dto.PbAds.ToAd(ad), nil
}

func (s *Server) DeleteAd(ctx context.Context, in *pb.DeleteAdRequest) (*empty.Empty, error) {
	if err := s.Service.DeleteAd(dto.PbAds.FromDeleteAdRequest(in)); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) ListPhotos(ctx context.Context, in *pb.ListPhotosRequest) (*pb.ListPhotosResponse, error) {
	adID := dto.PbAds.FromListPhotosRequest(in)

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
	return dto.PbAds.ToListPhotosResponse(result), nil
}

func (s *Server) ListTags(ctx context.Context, in *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	adID := dto.PbAds.FromListTagsRequest(in)

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
	return dto.PbAds.ToListTagsResponse(result), nil
}
