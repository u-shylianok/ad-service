package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/internal/testing/mocks/repository"
)

func TestPhotoService_ListAdPhotos(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){

		"valid result": func(t *testing.T) {

			expectedPhotos := []string{"photo 1", "photo 2"}
			expectedError := error(nil)

			repo := repository.PhotoMock{}
			repo.ListLinksByAdReturns(
				expectedPhotos,
				expectedError,
			)

			ps := &PhotoService{
				photoRepo: &repo,
			}

			photos, err := ps.ListAdPhotos(1)
			require.NoError(t, err)
			require.EqualValues(t, expectedPhotos, photos)
		},
		"error result": func(t *testing.T) {

			expectedError := fmt.Errorf("some error")

			repo := repository.PhotoMock{}
			repo.ListLinksByAdReturns(
				nil,
				expectedError,
			)

			ps := &PhotoService{
				photoRepo: &repo,
			}

			_, err := ps.ListAdPhotos(1)
			require.Error(t, err)
			require.EqualError(t, err, expectedError.Error())
		},
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
