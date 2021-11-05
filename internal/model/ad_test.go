package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdRequest_Validate(t *testing.T) {
	cases := []struct {
		name        string
		in          AdRequest
		expectError bool
		expected    error
	}{
		{
			name: "valid request without additional info",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "description",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
			},
			expectError: false,
		},
		{
			name: "valid request with photos info",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "description",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
				OtherPhotos: &[]string{
					"https://picsum.photos/id/102/200/200",
					"https://picsum.photos/id/103/200/200",
				},
			},
			expectError: false,
		},
		{
			name: "valid request with tags info",
			in: AdRequest{
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
			expectError: false,
		},
		{
			name: "invalid request (field Name is missing)",
			in: AdRequest{
				Price:       100,
				Description: "description",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
			},
			expectError: true,
			expected:    fmt.Errorf("name should not be empty"),
		},
		{
			name: "invalid request (field Price is missing)",
			in: AdRequest{
				Name:        "name",
				Description: "description",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
			},
			expectError: true,
			expected:    fmt.Errorf("price must be greater than zero"),
		},
		{
			name: "invalid request (field Description is missing)",
			in: AdRequest{
				Name:      "name",
				Price:     100,
				MainPhoto: "https://picsum.photos/id/101/200/200",
			},
			expectError: true,
			expected:    fmt.Errorf("description should not be empty"),
		},
		{
			name: "invalid request (field MainPhoto is missing)",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "description",
			},
			expectError: true,
			expected:    fmt.Errorf("main photo must exist"),
		},
		{
			name:        "invalid request (empty struct)",
			in:          AdRequest{},
			expectError: false,
			expected:    fmt.Errorf("name should not be empty"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.in.Validate()
			if tc.expectError {
				require.Error(t, err)
				require.EqualError(t, tc.expected, err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
