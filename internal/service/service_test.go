package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/internal/repository"
	"github.com/u-shylianok/ad-service/internal/secure"
	"github.com/u-shylianok/ad-service/internal/service"
	repoMock "github.com/u-shylianok/ad-service/internal/testing/mocks/repository"
	secureMock "github.com/u-shylianok/ad-service/internal/testing/mocks/secure"
)

func TestNewService(t *testing.T) {
	type args struct {
		repos  *repository.Repository
		secure *secure.Secure
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
					User:  &repoMock.UserMock{},
					Ad:    &repoMock.AdMock{},
					Photo: &repoMock.PhotoMock{},
					Tag:   &repoMock.TagMock{},
				},
				secure: &secure.Secure{
					Hasher: &secureMock.HasherMock{},
				},
			},
			want: &service.Service{
				Auth:  service.NewAuthService(&repoMock.UserMock{}, &secureMock.HasherMock{}),
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

			service := service.NewService(test.args.repos, test.args.secure)
			require.Equal(t, test.want, service)
		})
	}
}
