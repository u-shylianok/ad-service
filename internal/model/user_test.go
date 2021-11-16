package model

import (
	"reflect"
	"testing"
)

func TestUser_ToResponse(t *testing.T) {
	cases := []struct {
		name string
		in   User
		want UserResponse
	}{
		{
			name: "valid response from user with all fields",
			in: User{
				ID:       1,
				Name:     "name",
				Username: "username",
				Password: "password",
			},
			want: UserResponse{
				Name:     "name",
				Username: "username",
			},
		},
		{
			name: "valid response from user with only required fields",
			in: User{
				Name:     "name",
				Username: "username",
			},
			want: UserResponse{
				Name:     "name",
				Username: "username",
			},
		},
		{
			name: "valid response from empty user",
			in:   User{},
			want: UserResponse{},
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
