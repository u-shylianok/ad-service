package model

import (
	"fmt"
	"net/url"
	"reflect"
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
			expectError: true,
			expected:    fmt.Errorf("name should not be empty"),
		},
	}

	for _, tc := range cases {
		tc := tc
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

func TestListAdsSortingParamsFromURL(t *testing.T) {
	type args struct {
		values url.Values
	}
	cases := []struct {
		name string
		args args
		want []AdsSortingParam
	}{
		{
			name: "valid url sort parameters",
			args: args{
				values: url.Values{
					"sortby": []string{"name", "date", "price", "description"},
					"order":  []string{"asc", "dsc", "asc", "dsc"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "name",
					IsDesc: false,
				},
				{
					Field:  "date",
					IsDesc: true,
				},
				{
					Field:  "price",
					IsDesc: false,
				},
				{
					Field:  "description",
					IsDesc: true,
				},
			},
		},
		{
			name: "valid url sort parameters (some parameters capitalized)",
			args: args{
				values: url.Values{
					"sortby": []string{"nAme", "DATE", "pricE", "Description"},
					"order":  []string{"asc", "DSC", "Asc", "dsc"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "name",
					IsDesc: false,
				},
				{
					Field:  "date",
					IsDesc: true,
				},
				{
					Field:  "price",
					IsDesc: false,
				},
				{
					Field:  "description",
					IsDesc: true,
				},
			},
		},
		{
			name: "valid url sort parameter (only price)",
			args: args{
				values: url.Values{
					"sortby": []string{"price"},
					"order":  []string{"dsc"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "price",
					IsDesc: true,
				},
			},
		},
		{
			name: "invalid url sort parameters (order skipped)",
			args: args{
				values: url.Values{
					"sortby": []string{"name", "date", "price", "description"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "name",
					IsDesc: false,
				},
				{
					Field:  "date",
					IsDesc: false,
				},
				{
					Field:  "price",
					IsDesc: false,
				},
				{
					Field:  "description",
					IsDesc: false,
				},
			},
		},
		{
			name: "invalid url sort parameters (sortby skipped)",
			args: args{
				values: url.Values{
					"order": []string{"asc", "dsc", "asc", "dsc"},
				},
			},
			want: nil,
		},
		{
			name: "invalid url sort parameters (some order parameters skipped and setted by default)",
			args: args{
				values: url.Values{
					"sortby": []string{"description", "price", "date", "name"},
					"order":  []string{"dsc"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "description",
					IsDesc: true,
				},
				{
					Field:  "price",
					IsDesc: false,
				},
				{
					Field:  "date",
					IsDesc: false,
				},
				{
					Field:  "name",
					IsDesc: false,
				},
			},
		},
		{
			name: "invalid url sort parameters (some parameters are corrupted)",
			args: args{
				values: url.Values{
					"sortby": []string{"nnname", "date", "priice", "description"},
					"order":  []string{"asc", "dsc", "asc", "dsc"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "date",
					IsDesc: true,
				},
				{
					Field:  "description",
					IsDesc: true,
				},
			},
		},
		{
			name: "invalid url sort parameters (unexpected and expected sortby and order parameters)",
			args: args{
				values: url.Values{
					"sortby": []string{"name", "test", "invalid", "description"},
					"order":  []string{"asc", "test", "asc", "dsc"},
				},
			},
			want: []AdsSortingParam{
				{
					Field:  "name",
					IsDesc: false,
				},
				{
					Field:  "description",
					IsDesc: true,
				},
			},
		},
		{
			name: "invalid url sort parameters (only unexpected sortby and order parameters)",
			args: args{
				values: url.Values{
					"sortby": []string{"joke", "test", "DROP", "false"},
					"order":  []string{"true", "test", "asc", "dsc"},
				},
			},
			want: nil,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ListAdsSortingParamsFromURL(tc.args.values)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot: %#v,\nwant: %#v", got, tc.want)
			}

			// // another testcases processing example

			// got := ListAdsSortingParamsFromURL(tc.args.values)
			// equal := reflect.DeepEqual(got, tc.want)
			// require.True(t, equal)
		})
	}
}
