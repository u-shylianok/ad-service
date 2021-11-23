package service

import (
	"fmt"
	"testing"
	"time"

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

func TestAdService_ListAds(t *testing.T) {
	type fields struct {
		adRepo   *repository.AdMock
		userRepo *repository.UserMock
	}
	type args struct {
		params []model.AdsSortingParam
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, []model.AdResponse, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.ListReturns([]model.Ad{
					{
						ID:          1,
						UserID:      1,
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{
						ID:          2,
						UserID:      2,
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, nil)

				userRepo := repository.UserMock{}
				userRepo.ListInIDsReturns([]model.User{
					{
						ID:       1,
						Name:     "name",
						Username: "username",
					},
					{
						ID:       2,
						Name:     "name2",
						Username: "username2",
					},
				}, nil)

				f.adRepo = &adRepo
				f.userRepo = &userRepo
			},
			args: args{nil},
			assert: func(t *testing.T, f *fields, response []model.AdResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, []model.AdResponse{
					{

						ID: 1,
						User: model.UserResponse{
							Name:     "name",
							Username: "username",
						},
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{

						ID: 2,
						User: model.UserResponse{
							Name:     "name2",
							Username: "username2",
						},
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, response)
			},
		},
		{
			name: "fail - adRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.ListReturns(nil, fmt.Errorf("some error"))

				f.adRepo = &adRepo
			},
			args: args{nil},
			assert: func(t *testing.T, f *fields, response []model.AdResponse, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, []model.AdResponse(nil), response)
			},
		},
		{
			name: "fail - userRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.ListReturns([]model.Ad{
					{
						ID:          1,
						UserID:      1,
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{
						ID:          2,
						UserID:      2,
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, nil)

				userRepo := repository.UserMock{}
				userRepo.ListInIDsReturns(nil, fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.userRepo = &userRepo
			},
			args: args{nil},
			assert: func(t *testing.T, f *fields, response []model.AdResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, []model.AdResponse{
					{

						ID:          1,
						User:        model.UserResponse{},
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{

						ID:          2,
						User:        model.UserResponse{},
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, response)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			adService := NewAdService(fields.adRepo, fields.userRepo, nil, nil)
			response, err := adService.ListAds(test.args.params)

			test.assert(t, &fields, response, err)
		})
	}
}

func TestAdService_SearchAds(t *testing.T) {
	type fields struct {
		adRepo   *repository.AdMock
		userRepo *repository.UserMock
	}
	type args struct {
		params model.AdFilter
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, []model.AdResponse, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.ListWithFilterReturns([]model.Ad{
					{
						ID:          1,
						UserID:      1,
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{
						ID:          2,
						UserID:      2,
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, nil)

				userRepo := repository.UserMock{}
				userRepo.ListInIDsReturns([]model.User{
					{
						ID:       1,
						Name:     "name",
						Username: "username",
					},
					{
						ID:       2,
						Name:     "name2",
						Username: "username2",
					},
				}, nil)

				f.adRepo = &adRepo
				f.userRepo = &userRepo
			},
			args: args{model.AdFilter{}},
			assert: func(t *testing.T, f *fields, response []model.AdResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, []model.AdResponse{
					{

						ID: 1,
						User: model.UserResponse{
							Name:     "name",
							Username: "username",
						},
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{

						ID: 2,
						User: model.UserResponse{
							Name:     "name2",
							Username: "username2",
						},
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, response)
			},
		},
		{
			name: "fail - adRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.ListWithFilterReturns(nil, fmt.Errorf("some error"))

				f.adRepo = &adRepo
			},
			args: args{model.AdFilter{}},
			assert: func(t *testing.T, f *fields, response []model.AdResponse, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, []model.AdResponse(nil), response)
			},
		},
		{
			name: "fail - userRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.ListWithFilterReturns([]model.Ad{
					{
						ID:          1,
						UserID:      1,
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{
						ID:          2,
						UserID:      2,
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, nil)

				userRepo := repository.UserMock{}
				userRepo.ListInIDsReturns(nil, fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.userRepo = &userRepo
			},
			args: args{model.AdFilter{}},
			assert: func(t *testing.T, f *fields, response []model.AdResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, []model.AdResponse{
					{

						ID:          1,
						User:        model.UserResponse{},
						Name:        "name",
						Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
						Price:       100,
						Description: "description",
						MainPhoto:   "https://picsum.photos/id/101/200/200",
					},
					{

						ID:          2,
						User:        model.UserResponse{},
						Name:        "name2",
						Date:        time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
						Price:       200,
						Description: "description2",
						MainPhoto:   "https://picsum.photos/id/201/200/200",
					},
				}, response)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			adService := NewAdService(fields.adRepo, fields.userRepo, nil, nil)
			response, err := adService.SearchAds(test.args.params)

			test.assert(t, &fields, response, err)
		})
	}
}

func TestAdService_GetAd(t *testing.T) {
	type fields struct {
		adRepo    *repository.AdMock
		userRepo  *repository.UserMock
		photoRepo *repository.PhotoMock
		tagRepo   *repository.TagMock
	}
	type args struct {
		adID   int
		fields model.AdOptionalFieldsParam
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, model.AdResponse, error)
	}{
		{
			name: "success - without optional fields",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.GetReturns(model.Ad{
					ID:          1,
					UserID:      1,
					Name:        "name",
					Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
				}, nil)

				userRepo := repository.UserMock{}
				userRepo.GetByIDReturns(model.User{
					Name:     "name",
					Username: "username",
				}, nil)

				f.adRepo = &adRepo
				f.userRepo = &userRepo
			},
			args: args{
				adID:   1,
				fields: model.AdOptionalFieldsParam{},
			},
			assert: func(t *testing.T, f *fields, response model.AdResponse, err error) {
				require.NoError(t, err)

				expected := model.AdResponse{
					ID: 1,
					User: model.UserResponse{
						Name:     "name",
						Username: "username",
					},
					Name:        "name",
					Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
				}
				require.Equal(t, expected, response)
			},
		},
		{
			name: "success - with all optional fields",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.GetReturns(model.Ad{
					ID:          1,
					UserID:      1,
					Name:        "name",
					Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
					Price:       100,
					Description: "description",
					MainPhoto:   "https://picsum.photos/id/101/200/200",
				}, nil)

				userRepo := repository.UserMock{}
				userRepo.GetByIDReturns(model.User{
					Name:     "name",
					Username: "username",
				}, nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.ListLinksByAdReturns([]string{
					"https://picsum.photos/id/102/200/200",
					"https://picsum.photos/id/103/200/200",
				}, nil)

				tagRepo := repository.TagMock{}
				tagRepo.ListNamesByAdReturns([]string{
					"tag 1",
					"tag 2",
				}, nil)

				f.adRepo = &adRepo
				f.userRepo = &userRepo
				f.photoRepo = &photoRepo
				f.tagRepo = &tagRepo
			},
			args: args{
				adID: 1,
				fields: model.AdOptionalFieldsParam{
					Photos: true,
					Tags:   true,
				},
			},
			assert: func(t *testing.T, f *fields, response model.AdResponse, err error) {
				require.NoError(t, err)

				expected := model.AdResponse{
					ID: 1,
					User: model.UserResponse{
						Name:     "name",
						Username: "username",
					},
					Name:        "name",
					Date:        time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
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
				}
				require.Equal(t, expected, response)
			},
		},
		{
			name: "fail - adRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.GetReturns(model.Ad{}, fmt.Errorf("some error"))

				f.adRepo = &adRepo
			},
			args: args{
				adID: 1,
			},
			assert: func(t *testing.T, f *fields, response model.AdResponse, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, model.AdResponse{}, response)
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
			response, err := adService.GetAd(test.args.adID, test.args.fields)

			test.assert(t, &fields, response, err)
		})
	}
}

func TestAdService_UpdateAd(t *testing.T) {
	type fields struct {
		adRepo    *repository.AdMock
		photoRepo *repository.PhotoMock
		tagRepo   *repository.TagMock
	}
	type args struct {
		userID int
		adID   int
		ad     model.AdRequest
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.UpdateReturns(nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.DeleteAllByAdReturns(nil)
				photoRepo.CreateListReturns(nil)

				tagRepo := repository.TagMock{}
				tagRepo.DetachAllFromAdReturns(nil)
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(0, 1, nil)
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(1, 2, nil)
				tagRepo.AttachToAdReturns(nil)

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
				f.tagRepo = &tagRepo
			},
			args: args{
				userID: 1,
				adID:   1,
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
			assert: func(t *testing.T, f *fields, err error) {
				require.NoError(t, err)

				require.Equal(t, 1, f.adRepo.UpdateCallCount())
				require.Equal(t, 1, f.photoRepo.DeleteAllByAdCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
				require.Equal(t, 1, f.tagRepo.DetachAllFromAdCallCount())
				require.Equal(t, 2, f.tagRepo.GetIDOrCreateIfNotExistsCallCount())
				require.Equal(t, 2, f.tagRepo.AttachToAdCallCount())
			},
		},
		{
			name: "fail - adRepo returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.UpdateReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
			},
			args: args{
				userID: 1,
				adID:   1,
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
			assert: func(t *testing.T, f *fields, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")

				require.Equal(t, 1, f.adRepo.UpdateCallCount())
			},
		},
		{
			name: "fail - photoRepo DeleteAllByAd returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.UpdateReturns(nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.DeleteAllByAdReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
			},
			args: args{
				userID: 1,
				adID:   1,
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
			assert: func(t *testing.T, f *fields, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")

				require.Equal(t, 1, f.adRepo.UpdateCallCount())
				require.Equal(t, 1, f.photoRepo.DeleteAllByAdCallCount())
			},
		},
		{
			name: "fail - photoRepo CreateList returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.UpdateReturns(nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.DeleteAllByAdReturns(nil)
				photoRepo.CreateListReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
			},
			args: args{
				userID: 1,
				adID:   1,
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
			assert: func(t *testing.T, f *fields, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")

				require.Equal(t, 1, f.adRepo.UpdateCallCount())
				require.Equal(t, 1, f.photoRepo.DeleteAllByAdCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
			},
		},
		{
			name: "failt - tagRepo DetachAllFromAd returns error",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.UpdateReturns(nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.DeleteAllByAdReturns(nil)
				photoRepo.CreateListReturns(nil)

				tagRepo := repository.TagMock{}
				tagRepo.DetachAllFromAdReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
				f.tagRepo = &tagRepo
			},
			args: args{
				userID: 1,
				adID:   1,
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
			assert: func(t *testing.T, f *fields, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")

				require.Equal(t, 1, f.adRepo.UpdateCallCount())
				require.Equal(t, 1, f.photoRepo.DeleteAllByAdCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
				require.Equal(t, 1, f.tagRepo.DetachAllFromAdCallCount())
			},
		},
		{
			name: "fail - tagRepo GetIDOrCreateIfNotExists returns error and one of attach failed",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.UpdateReturns(nil)

				photoRepo := repository.PhotoMock{}
				photoRepo.DeleteAllByAdReturns(nil)
				photoRepo.CreateListReturns(nil)

				tagRepo := repository.TagMock{}
				tagRepo.DetachAllFromAdReturns(nil)
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(0, 1, fmt.Errorf("some error"))
				tagRepo.GetIDOrCreateIfNotExistsReturnsOnCall(1, 2, nil)
				tagRepo.AttachToAdReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
				f.photoRepo = &photoRepo
				f.tagRepo = &tagRepo
			},
			args: args{
				userID: 1,
				adID:   1,
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
			assert: func(t *testing.T, f *fields, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")

				require.Equal(t, 1, f.adRepo.UpdateCallCount())
				require.Equal(t, 1, f.photoRepo.DeleteAllByAdCallCount())
				require.Equal(t, 1, f.photoRepo.CreateListCallCount())
				require.Equal(t, 1, f.tagRepo.DetachAllFromAdCallCount())
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

			adService := NewAdService(fields.adRepo, nil, fields.photoRepo, fields.tagRepo)
			err := adService.UpdateAd(test.args.userID, test.args.adID, test.args.ad)

			test.assert(t, &fields, err)
		})
	}
}

func TestAdService_DeleteAd(t *testing.T) {
	type fields struct {
		adRepo *repository.AdMock
	}
	type args struct {
		userID int
		adID   int
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.DeleteReturns(nil)

				f.adRepo = &adRepo
			},
			args: args{
				userID: 1,
				adID:   1,
			},
			assert: func(t *testing.T, f *fields, err error) {
				require.NoError(t, err)

				require.Equal(t, 1, f.adRepo.DeleteCallCount())
			},
		},
		{
			name: "success - some errors occurred",
			setup: func(f *fields) {
				adRepo := repository.AdMock{}
				adRepo.DeleteReturns(fmt.Errorf("some error"))

				f.adRepo = &adRepo
			},
			args: args{
				userID: 1,
				adID:   1,
			},
			assert: func(t *testing.T, f *fields, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")

				require.Equal(t, 1, f.adRepo.DeleteCallCount())
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			adService := NewAdService(fields.adRepo, nil, nil, nil)
			err := adService.DeleteAd(test.args.userID, test.args.adID)

			test.assert(t, &fields, err)
		})
	}
}
