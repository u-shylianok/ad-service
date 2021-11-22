package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/internal/testing/mocks/repository"
)

func TestPhotoService_ListAdPhotos(t *testing.T) {
	type fields struct {
		photoRepo *repository.PhotoMock
	}
	type args struct {
		adID int
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, []string, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				repo := &repository.PhotoMock{}
				repo.ListLinksByAdReturns([]string{"photo 1", "photo 2"}, nil)

				f.photoRepo = repo
			},
			args: args{
				adID: 1,
			},
			assert: func(t *testing.T, f *fields, photos []string, err error) {
				require.NoError(t, err)
				require.Equal(t, []string{"photo 1", "photo 2"}, photos)
			},
		},
		{
			name: "fail - some errors occurred",
			setup: func(f *fields) {
				repo := &repository.PhotoMock{}
				repo.ListLinksByAdReturns(nil, fmt.Errorf("some error"))

				f.photoRepo = repo
			},
			args: args{
				adID: 1,
			},
			assert: func(t *testing.T, f *fields, photos []string, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, []string(nil), photos)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			photoService := NewPhotoService(fields.photoRepo)
			photos, err := photoService.ListAdPhotos(test.args.adID)

			test.assert(t, &fields, photos, err)
		})
	}
}

func TestPhotoService_ListPhotos(t *testing.T) {
	type fields struct {
		photoRepo *repository.PhotoMock
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		assert func(*testing.T, *fields, []string, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				repo := &repository.PhotoMock{}
				repo.ListLinksReturns([]string{"photo 1", "photo 2", "photo 3", "photo 4"}, nil)

				f.photoRepo = repo
			},
			assert: func(t *testing.T, f *fields, photos []string, err error) {
				require.NoError(t, err)
				require.Equal(t, []string{"photo 1", "photo 2", "photo 3", "photo 4"}, photos)
			},
		},
		{
			name: "fail - some errors occurred",
			setup: func(f *fields) {
				repo := &repository.PhotoMock{}
				repo.ListLinksReturns(nil, fmt.Errorf("some error"))

				f.photoRepo = repo
			},
			assert: func(t *testing.T, f *fields, photos []string, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, []string(nil), photos)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			photoService := NewPhotoService(fields.photoRepo)
			photos, err := photoService.ListPhotos()

			test.assert(t, &fields, photos, err)
		})
	}
}
