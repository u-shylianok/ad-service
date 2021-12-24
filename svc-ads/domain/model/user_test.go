package model_test

import (
	"reflect"
	"testing"

	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
)

func TestUser_ToResponse(t *testing.T) {
	cases := []struct {
		name string
		in   model.User
		want model.UserResponse
	}{
		{
			name: "valid response from user with all fields",
			in: model.User{
				ID:       1,
				Name:     "name",
				Username: "username",
				Password: "password",
			},
			want: model.UserResponse{
				Name:     "name",
				Username: "username",
			},
		},
		{
			name: "valid response from user with only required fields",
			in: model.User{
				Name:     "name",
				Username: "username",
			},
			want: model.UserResponse{
				Name:     "name",
				Username: "username",
			},
		},
		{
			name: "valid response from empty user",
			in:   model.User{},
			want: model.UserResponse{},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.ToResponse()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot: %#v,\nwant: %#v", got, tc.want)
			}
		})
	}
}
