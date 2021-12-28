package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/service"
)

type Server struct {
	pb.UnimplementedAdServiceServer

	Service service.Service
}

func (s *Server) ListAds(ctx context.Context, req *pb.ListAdsRequest) (*pb.ListAdsResponse, error) {
	s.Service.Ad.ListAds(nil)
	return nil, nil
}

func (s *Server) SearchAds(context.Context, *pb.SearchAdsRequest) (*pb.SearchAdsResponse, error) {
	return nil, nil
}

func (s *Server) GetAd(context.Context, *pb.GetAdRequest) (*pb.Ad, error) {
	return nil, nil
}

func (s *Server) CreateAd(context.Context, *pb.CreateAdRequest) (*pb.Ad, error) {
	return nil, nil
}

func (s *Server) UpdateAd(context.Context, *pb.UpdateAdRequest) (*pb.Ad, error) {
	return nil, nil
}

func (s *Server) DeleteAd(context.Context, *pb.DeleteAdRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *Server) ListPhotos(context.Context, *pb.ListPhotosRequest) (*pb.ListPhotosResponse, error) {
	return nil, nil
}

func (s *Server) ListTags(context.Context, *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	return nil, nil
}
