package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/testing/mocks/repository"
)

func TestAdService_CreateAd(t *testing.T) {
	type args struct {
		userID int
		ad     model.AdRequest
	}
	type fields struct {
		adRepo    *repository.AdMock
		userRepo  *repository.UserMock
		photoRepo *repository.PhotoMock
		tagRepo   *repository.TagMock
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, int, error)
	}{
		{
			name: "success - without optional fields",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.CreateReturns(10, nil)

				f.adRepo = &adRepo
			},
			args: args{
				userID: 3,
				ad: model.AdRequest{
					Name:        "name",
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
				},
			},
			assert: func(t *testing.T, f *fields, adID int, err error) {
				require.NoError(t, err)
				require.EqualValues(t, 10, adID)

				require.Equal(t, 1, f.adRepo.CreateCallCount())
			},
		},
		{
			name: "success - with all fields",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.CreateReturns(10, nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.CreateListReturns(nil)

				tagRepo := repository.TagMock{}
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(0, 101, nil)
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(1, 102, nil)
				tagRepo.AttachToAdReturns(nil)

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
				f.tagRepo = &tagRepo
			},
			args: args{
				userID: 3,
				ad: model.AdRequest{
					Name:        "name",
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
					OtherPhotos: &[]string{
						"https://picsum.photos/id/102/200/200",
						"https://picsum.photos/id/103/200/200",
					},
					Tags: &[]string{
						"tag 1",
						"tag 2",
					},
				},
			},
			assert: func(t *testing.T, f *fields, adID int, err error) {
				require.NoError(t, err)
				require.EqualValues(t, 10, adID)

				require.Equal(t, 1, f.adRepo.CreateCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
				require.Equal(t, 2, f.tagRepo.GetIDOrCreateIfNotExistsCallCount())
				require.Equal(t, 2, f.tagRepo.AttachToAdCallCount())
			},
		},
		{
			name: "fail - adRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.CreateReturns(0, fmt.Errorf("some error"))

				f.adRepo = &adRepo
			},
			args: args{
				userID: 3,
				ad: model.AdRequest{
					Name:        "name",
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
				},
			},
			assert: func(t *testing.T, f *fields, adID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.EqualValues(t, 0, adID)

				require.Equal(t, 1, f.adRepo.CreateCallCount())
			},
		},

		{
			name: "fail - photoRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.CreateReturns(10, nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.CreateListReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
			},
			args: args{
				userID: 3,
				ad: model.AdRequest{
					Name:        "name",
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
					OtherPhotos: &[]string{
						"https://picsum.photos/id/102/200/200",
						"https://picsum.photos/id/103/200/200",
					},
				},
			},
			assert: func(t *testing.T, f *fields, adID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.EqualValues(t, 10, adID)

				require.Equal(t, 1, f.adRepo.CreateCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
			},
		},
		{
			name: "fail - tagRepo returns error and tagID next time, but attach fails",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.CreateReturns(10, nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.CreateListReturns(nil)

				tagRepo := repository.TagMock{}
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(0, 0, fmt.Errorf("some error 1"))
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(1, 102, nil)
				tagRepo.AttachToAdReturnsOnCall(0, fmt.Errorf("some error 2"))

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
				f.tagRepo = &tagRepo
			},
			args: args{
				userID: 3,
				ad: model.AdRequest{
					Name:        "name",
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
					OtherPhotos: &[]string{
						"https://picsum.photos/id/102/200/200",
						"https://picsum.photos/id/103/200/200",
					},
					Tags: &[]string{
						"tag 1",
						"tag 2",
					},
				},
			},
			assert: func(t *testing.T, f *fields, adID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error 2")
				require.EqualValues(t, 10, adID)

				require.Equal(t, 1, f.adRepo.CreateCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
				require.Equal(t, 2, f.tagRepo.GetIDOrCreateIfNotExistsCallCount())
				require.Equal(t, 1, f.tagRepo.AttachToAdCallCount())
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			adService := NewAdService(fields.adRepo, fields.userRepo, fields.photoRepo, fields.tagRepo)
			adID, err := adService.CreateAd(test.args.userID, test.args.ad)

			test.assert(t, &fields, adID, err)
		})
	}
}
