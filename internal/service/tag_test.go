package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/internal/testing/mocks/repository"
)

func TestTagService_ListAdTags(t *testing.T) {
	type fields struct {
		tagRepo *repository.TagMock
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
				repo := &repository.TagMock{}
				repo.ListNamesByAdReturns([]string{"tag 1", "tag 2"}, nil)

				f.tagRepo = repo
			},
			args: args{
				adID: 1,
			},
			assert: func(t *testing.T, f *fields, tags []string, err error) {
				require.NoError(t, err)
				require.Equal(t, []string{"tag 1", "tag 2"}, tags)
			},
		},
		{
			name: "fail - some errors occurred",
			setup: func(f *fields) {
				repo := &repository.TagMock{}
				repo.ListNamesByAdReturns(nil, fmt.Errorf("some error"))

				f.tagRepo = repo
			},
			args: args{
				adID: 1,
			},
			assert: func(t *testing.T, f *fields, tags []string, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, []string(nil), tags)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			tagService := NewTagService(fields.tagRepo)
			tags, err := tagService.ListAdTags(test.args.adID)

			test.assert(t, &fields, tags, err)
		})
	}
}

func TestTagService_ListTags(t *testing.T) {
	type fields struct {
		tagRepo *repository.TagMock
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		assert func(*testing.T, *fields, []string, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				repo := &repository.TagMock{}
				repo.ListNamesReturns([]string{"tag 1", "tag 2", "tag 3", "tag 4"}, nil)

				f.tagRepo = repo
			},
			assert: func(t *testing.T, f *fields, tags []string, err error) {
				require.NoError(t, err)
				require.Equal(t, []string{"tag 1", "tag 2", "tag 3", "tag 4"}, tags)
			},
		},
		{
			name: "fail - some errors occurred",
			setup: func(f *fields) {
				repo := &repository.TagMock{}
				repo.ListNamesReturns(nil, fmt.Errorf("some error"))

				f.tagRepo = repo
			},
			assert: func(t *testing.T, f *fields, tags []string, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, []string(nil), tags)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			tagService := NewTagService(fields.tagRepo)
			tags, err := tagService.ListTags()

			test.assert(t, &fields, tags, err)
		})
	}
}
