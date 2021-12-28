package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
	"github.com/u-shylianok/ad-service/svc-ads/service"
	repoMock "github.com/u-shylianok/ad-service/svc-ads/testing/mocks/repository"
)

func TestNewService(t *testing.T) {
	type args struct {
		repos *repository.Repository
	}
	tests := []struct {
		name string
		args args
		want *service.Service
	}{
		{
			name: "success",
			args: args{
				repos: &repository.Repository{
					Ad:    &repoMock.AdMock{},
					Photo: &repoMock.PhotoMock{},
					Tag:   &repoMock.TagMock{},
				},
			},
			want: &service.Service{
				Ad:    service.NewAdService(&repoMock.AdMock{}, &repoMock.UserMock{}, &repoMock.PhotoMock{}, &repoMock.TagMock{}),
				Photo: service.NewPhotoService(&repoMock.PhotoMock{}),
				Tag:   service.NewTagService(&repoMock.TagMock{}),
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			service := service.NewService(test.args.repos)
			require.Equal(t, test.want, service)
		})
	}
}
