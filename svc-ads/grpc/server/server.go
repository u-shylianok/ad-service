package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	pb "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/grpc/client"
	"github.com/u-shylianok/ad-service/svc-ads/model"
	"github.com/u-shylianok/ad-service/svc-ads/service"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Server) ListAds(ctx context.Context, in *pb.ListAdsRequest) (*pb.ListAdsResponse, error) {
	logrus.WithField("in", in).Error("ListAds")
	return nil, nil
}

func (s *Server) SearchAds(ctx context.Context, in *pb.SearchAdsRequest) (*pb.SearchAdsResponse, error) {
	logrus.WithField("in", in).Error("SearchAds")
	return nil, nil
}

func (s *Server) CreateAd(ctx context.Context, in *pb.CreateAdRequest) (*pb.Ad, error) {
	logrus.WithField("in", in).Error("CreateAd")
	return nil, nil
}

func (s *Server) UpdateAd(ctx context.Context, in *pb.UpdateAdRequest) (*pb.Ad, error) {
	logrus.WithField("in", in).Error("UpdateAd")
	return nil, nil
}

func (s *Server) DeleteAd(ctx context.Context, in *pb.DeleteAdRequest) (*empty.Empty, error) {
	logrus.WithField("in", in).Error("DeleteAd")
	return nil, nil
}

func (s *Server) ListPhotos(ctx context.Context, in *pb.ListPhotosRequest) (*pb.ListPhotosResponse, error) {
	logrus.WithField("in", in).Error("ListPhotos")
	return nil, nil
}

func (s *Server) ListTags(ctx context.Context, in *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	logrus.WithField("in", in).Error("ListTags")
	return nil, nil
}
