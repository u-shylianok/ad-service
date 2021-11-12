package model

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
	"time"

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
			name: "invalid request (got empty OtherPhoto)",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "description",
				OtherPhotos: &[]string{""},
			},
			expectError: true,
			expected:    fmt.Errorf("photo link should not be empty"),
		},
		{
			name: "invalid request (got empty Tag)",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "description",
				Tags:        &[]string{""},
			},
			expectError: true,
			expected:    fmt.Errorf("tag name should not be empty"),
		},
		{
			name: "invalid request (more than 200 symbols name)",
			in: AdRequest{
				Name:        "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
				Price:       100,
				Description: "description",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
			},
			expectError: true,
			expected:    fmt.Errorf("name should be no more than 200 symbols"),
		},
		{
			name: "invalid request (more than 1000 symbols description)",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
			},
			expectError: true,
			expected:    fmt.Errorf("description should be no more than 1000 symbols"),
		},
		{
			name: "invalid request (more than 3 photos)",
			in: AdRequest{
				Name:        "name",
				Price:       100,
				Description: "description",
				MainPhoto:   "https://picsum.photos/id/101/200/200",
				OtherPhotos: &[]string{
					"https://picsum.photos/id/102/200/200",
					"https://picsum.photos/id/103/200/200",
					"https://picsum.photos/id/104/200/200",
				},
			},
			expectError: true,
			expected:    fmt.Errorf("should be no more than 2 other photos"),
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
			name: "valid url sort parameter (empty values)",
			args: args{
				values: url.Values{},
			},
			want: nil,
		},
		{
			name: "valid url sort parameter (nil values)",
			args: args{
				values: nil,
			},
			want: nil,
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
		})
	}
}

func TestGetAdOptionalFieldsFromURL(t *testing.T) {
	type args struct {
		values url.Values
	}
	cases := []struct {
		name string
		args args
		want AdOptionalFieldsParam
	}{
		{
			name: "valid all optional fields",
			args: args{
				values: url.Values{
					"fields": []string{"description", "photos", "tags"},
				},
			},
			want: AdOptionalFieldsParam{
				Description: true,
				Photos:      true,
				Tags:        true,
			},
		},
		{
			name: "valid optional fields (some parameters capitalized)",
			args: args{
				values: url.Values{
					"fields": []string{"Description", "phOtOs", "TAGS"},
				},
			},
			want: AdOptionalFieldsParam{
				Description: true,
				Photos:      true,
				Tags:        true,
			},
		},
		{
			name: "valid description optional field",
			args: args{
				values: url.Values{
					"fields": []string{"description"},
				},
			},
			want: AdOptionalFieldsParam{
				Description: true,
			},
		},
		{
			name: "valid photos optional field",
			args: args{
				values: url.Values{
					"fields": []string{"photos"},
				},
			},
			want: AdOptionalFieldsParam{
				Photos: true,
			},
		},
		{
			name: "valid tags optional field",
			args: args{
				values: url.Values{
					"fields": []string{"tags"},
				},
			},
			want: AdOptionalFieldsParam{
				Tags: true,
			},
		},
		{
			name: "valid optional fields (empty fields)",
			args: args{
				values: url.Values{},
			},
			want: AdOptionalFieldsParam{},
		},
		{
			name: "valid optional fields (nil values)",
			args: args{
				values: nil,
			},
			want: AdOptionalFieldsParam{},
		},
		{
			name: "invalid optional fields (some parameters are corrupted)",
			args: args{
				values: url.Values{
					"fields": []string{"tags", "joke", "test", "description", "pphotooss"},
				},
			},
			want: AdOptionalFieldsParam{
				Description: true,
				Tags:        true,
			},
		},
		{
			name: "invalid optional fields (unexpected parameters only)",
			args: args{
				values: url.Values{
					"fields": []string{"name", "price", "test", "frog"},
				},
			},
			want: AdOptionalFieldsParam{},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := GetAdOptionalFieldsFromURL(tc.args.values)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot: %#v,\nwant: %#v", got, tc.want)
			}
		})
	}
}

func TestGetAdFilterFromURL(t *testing.T) {
	type args struct {
		values url.Values
	}
	cases := []struct {
		name string
		args args
		want AdFilter
	}{
		{
			name: "valid filter params",
			args: args{
				values: url.Values{
					"username":  []string{"test"},
					"startdate": []string{"2021-10-12"},
					"enddate":   []string{"2021-10-15"},
					"tags":      []string{"ТЕСТ", "КРАСНЫЙ"},
				},
			},
			want: AdFilter{
				Username:  "test",
				StartDate: time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2021, 10, 15, 0, 0, 0, 0, time.UTC),
				Tags:      []string{"ТЕСТ", "КРАСНЫЙ"},
			},
		},
		{
			name: "valid filter params without startdate",
			args: args{
				values: url.Values{
					"username": []string{"test"},
					"enddate":  []string{"2021-10-15"},
					"tags":     []string{"ТЕСТ", "КРАСНЫЙ"},
				},
			},
			want: AdFilter{
				Username: "test",
				EndDate:  time.Date(2021, 10, 15, 0, 0, 0, 0, time.UTC),
				Tags:     []string{"ТЕСТ", "КРАСНЫЙ"},
			},
		},
		{
			name: "valid filter params without enddate",
			args: args{
				values: url.Values{
					"username":  []string{"test"},
					"startdate": []string{"2021-10-12"},
					"tags":      []string{"ТЕСТ", "КРАСНЫЙ"},
				},
			},
			want: AdFilter{
				Username:  "test",
				StartDate: time.Date(2021, 10, 12, 0, 0, 0, 0, time.UTC),
				Tags:      []string{"ТЕСТ", "КРАСНЫЙ"},
			},
		},
		{
			name: "valid filter params (only username)",
			args: args{
				values: url.Values{
					"username": []string{"test"},
				},
			},
			want: AdFilter{
				Username: "test",
			},
		},
		{
			name: "valid filter params (only tags)",
			args: args{
				values: url.Values{
					"tags": []string{"ТЕСТ", "КРАСНЫЙ"},
				},
			},
			want: AdFilter{
				Tags: []string{"ТЕСТ", "КРАСНЫЙ"},
			},
		},
		{
			name: "valid filter params (empty params)",
			args: args{
				values: url.Values{},
			},
			want: AdFilter{},
		},
		{
			name: "valid optional fields (nil values)",
			args: args{
				values: nil,
			},
			want: AdFilter{},
		},
		{
			name: "invalid filter params (dates format is corrupted)",
			args: args{
				values: url.Values{
					"username":  []string{"test"},
					"startdate": []string{"2021/10/12"},
					"enddate":   []string{"2021/10/15"},
					"tags":      []string{"ТЕСТ", "КРАСНЫЙ"},
				},
			},
			want: AdFilter{
				Username: "test",
				Tags:     []string{"ТЕСТ", "КРАСНЫЙ"},
			},
		},
		{
			name: "invalid filter params (unexpected params)",
			args: args{
				values: url.Values{
					"name": []string{"test"},
					"date": []string{"2021/10/12"},
					"tag":  []string{"ТЕСТ"},
				},
			},
			want: AdFilter{},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := GetAdFilterFromURL(tc.args.values)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot: %#v,\nwant: %#v", got, tc.want)
			}
		})
	}
}
