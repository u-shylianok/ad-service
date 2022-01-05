package server

import (
	"context"

	pb "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/model"
	"github.com/u-shylianok/ad-service/svc-ads/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedAdServiceServer

	Service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{Service: service}
}

// func (s *Server) ListAds(ctx context.Context, req *pb.ListAdsRequest) (*pb.ListAdsResponse, error) {
// 	s.Service.Ad.ListAds(nil)
// 	return nil, nil
// }

// func (s *Server) SearchAds(context.Context, *pb.SearchAdsRequest) (*pb.SearchAdsResponse, error) {
// 	return nil, nil
// }

func (s *Server) GetAd(ctx context.Context, in *pb.GetAdRequest) (*pb.Ad, error) {
	ad, err := s.Service.GetAd(in.GetId(), model.AdOptionalFieldsParam{})
	if err != nil {
		return nil, err
	}
	out := &pb.Ad{
		Id:          ad.ID,
		Name:        ad.Name,
		Date:        timestamppb.New(ad.Date),
		Price:       int32(ad.Price),
		Description: ad.Description,
		Photo:       ad.MainPhoto,
		Photos:      make([]string, 0),
		Tags:        make([]string, 0),
	}
	return out, nil
}

// func (s *Server) CreateAd(context.Context, *pb.CreateAdRequest) (*pb.Ad, error) {
// 	return nil, nil
// }

// func (s *Server) UpdateAd(context.Context, *pb.UpdateAdRequest) (*pb.Ad, error) {
// 	return nil, nil
// }

// func (s *Server) DeleteAd(context.Context, *pb.DeleteAdRequest) (*empty.Empty, error) {
// 	return nil, nil
// }

// func (s *Server) ListPhotos(context.Context, *pb.ListPhotosRequest) (*pb.ListPhotosResponse, error) {
// 	return nil, nil
// }

// func (s *Server) ListTags(context.Context, *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
// 	return nil, nil
// }
