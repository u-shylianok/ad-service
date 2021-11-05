package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdRequest_Validate(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"valid":                      testValidRequest,
		"valid with photos":          testValidRequestWithPhotos,
		"valid with tags":            testValidRequestWithTags,
		"valid with photos and tags": testValidRequestWithPhotosAndTags,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testValidRequest(t *testing.T) {
	adRequest := AdRequest{
		Name:        "ad name",
		Price:       100,
		Description: "test ad",
		MainPhoto:   "https://picsum.photos/id/101/200/200",
	}

	err := adRequest.Validate()
	require.NoError(t, err)
}

func testValidRequestWithPhotos(t *testing.T) {
	adRequest := AdRequest{
		Name:        "ad name",
		Price:       100,
		Description: "test ad",
		MainPhoto:   "https://picsum.photos/id/101/200/200",
		OtherPhotos: &[]string{
			"https://picsum.photos/id/102/200/200",
			"https://picsum.photos/id/103/200/200",
		},
	}

	err := adRequest.Validate()
	require.NoError(t, err)
}

func testValidRequestWithTags(t *testing.T) {
	adRequest := AdRequest{
		Name:        "ad name",
		Price:       100,
		Description: "test ad",
		MainPhoto:   "https://picsum.photos/id/101/200/200",
		Tags: &[]string{
			"tag 1",
			"tag 2",
		},
	}

	err := adRequest.Validate()
	require.NoError(t, err)
}

func testValidRequestWithPhotosAndTags(t *testing.T) {
	adRequest := AdRequest{
		Name:        "ad name",
		Price:       100,
		Description: "test ad",
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

	err := adRequest.Validate()
	require.NoError(t, err)
}

// func TestAdRequest_Validate(t *testing.T) {
// 	type fields struct {
// 		Name        string
// 		Price       int
// 		Description string
// 		MainPhoto   string
// 		OtherPhotos *[]string
// 		Tags        *[]string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := AdRequest{
// 				Name:        tt.fields.Name,
// 				Price:       tt.fields.Price,
// 				Description: tt.fields.Description,
// 				MainPhoto:   tt.fields.MainPhoto,
// 				OtherPhotos: tt.fields.OtherPhotos,
// 				Tags:        tt.fields.Tags,
// 			}
// 			if err := r.Validate(); (err != nil) != tt.wantErr {
// 				t.Errorf("AdRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
