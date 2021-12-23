package service_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/ad-service/domain/model"
	"github.com/u-shylianok/ad-service/internal/service"
	"github.com/u-shylianok/ad-service/internal/testing/mocks/repository"
	"github.com/u-shylianok/ad-service/internal/testing/mocks/secure"
)

func TestAuthService_CreateUser(t *testing.T) {
	type fields struct {
		repo   *repository.UserMock
		hasher *secure.HasherMock
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, int, error)
	}{
		{
			name: "success - create new user without errors",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, sql.ErrNoRows)
				userRepo.CreateReturns(1, nil)

				hasher := secure.HasherMock{}
				hasher.HashPasswordReturns("some hash", nil)

				f.repo = &userRepo
				f.hasher = &hasher
			},
			args: args{
				user: model.User{
					Name:     "name",
					Username: "username",
					Password: "password",
				},
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.NoError(t, err)
				require.Equal(t, 1, userID)

				require.Equal(t, 1, f.repo.GetCallCount())
				require.Equal(t, 1, f.repo.CreateCallCount())

				user := f.repo.CreateArgsForCall(0)
				require.Equal(t, "some hash", user.Password)
			},
		},
		{
			name: "fail - username is invalid or already taken",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, nil)

				f.repo = &userRepo
			},
			args: args{
				user: model.User{
					Name:     "name",
					Username: "username",
					Password: "password",
				},
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "username is invalid or already taken")
				require.Equal(t, 0, userID)
			},
		},
		{
			name: "fail - userRepo Get returns error",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, fmt.Errorf("some error"))

				f.repo = &userRepo
			},
			args: args{
				user: model.User{
					Name:     "name",
					Username: "username",
					Password: "password",
				},
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, 0, userID)
			},
		},
		{
			name: "fail - hashPassword returns error",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, sql.ErrNoRows)

				hasher := secure.HasherMock{}
				hasher.HashPasswordReturns("", fmt.Errorf("some error"))

				f.repo = &userRepo
				f.hasher = &hasher
			},
			args: args{
				user: model.User{
					Name:     "name",
					Username: "username",
					Password: "password",
				},
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, 0, userID)
			},
		},
		{
			name: "fail - userRepo Create returns error",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, sql.ErrNoRows)
				userRepo.CreateReturns(0, fmt.Errorf("some error"))

				hasher := secure.HasherMock{}
				hasher.HashPasswordReturns("some hash", nil)

				f.repo = &userRepo
				f.hasher = &hasher
			},
			args: args{
				user: model.User{
					Name:     "name",
					Username: "username",
					Password: "password",
				},
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, 0, userID)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			authService := service.NewAuthService(fields.repo, fields.hasher)
			userID, err := authService.CreateUser(test.args.user)

			test.assert(t, &fields, userID, err)
		})
	}
}

func TestAuthService_CheckUser(t *testing.T) {
	type fields struct {
		repo   *repository.UserMock
		hasher *secure.HasherMock
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, int, error)
	}{
		{
			name: "success - without errors",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{
					ID:       1,
					Name:     "name",
					Username: "username",
					Password: "some hash",
				}, nil)

				hasher := secure.HasherMock{}
				hasher.CheckPasswordHashReturns(true)

				f.repo = &userRepo
				f.hasher = &hasher
			},
			args: args{
				username: "username",
				password: "password",
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.NoError(t, err)
				require.Equal(t, 1, userID)

				password, hash := f.hasher.CheckPasswordHashArgsForCall(0)
				require.Equal(t, "password", password)
				require.Equal(t, "some hash", hash)
			},
		},
		{
			name: "fail - incorrect username",
			setup: func(f *fields) {

				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, sql.ErrNoRows)

				f.repo = &userRepo
			},
			args: args{
				username: "username",
				password: "password",
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "incorrect username or password")
				require.Equal(t, 0, userID)
			},
		},
		{
			name: "fail - incorrect password",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{
					ID:       1,
					Name:     "name",
					Username: "username",
					Password: "some hash",
				}, nil)

				hasher := secure.HasherMock{}
				hasher.CheckPasswordHashReturns(false)

				f.repo = &userRepo
				f.hasher = &hasher
			},
			args: args{
				username: "username",
				password: "password",
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "incorrect username or password")
				require.Equal(t, 0, userID)
			},
		},

		{
			name: "fail - userRepo Get returns error",
			setup: func(f *fields) {
				userRepo := repository.UserMock{}
				userRepo.GetReturns(model.User{}, fmt.Errorf("some error"))

				f.repo = &userRepo
			},
			args: args{
				username: "username",
				password: "password",
			},
			assert: func(t *testing.T, f *fields, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "some error")
				require.Equal(t, 0, userID)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			authService := service.NewAuthService(fields.repo, fields.hasher)
			userID, err := authService.CheckUser(test.args.username, test.args.password)

			test.assert(t, &fields, userID, err)
		})
	}
}

func TestAuthService_GenerateToken(t *testing.T) {
	const defaultTokenTTL = 12 * time.Hour

	type fields struct {
		time int64
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name   string
		setup  func(*fields)
		args   args
		assert func(*testing.T, *fields, *service.AuthService, string, int64, error)
	}{
		{
			name: "success",
			setup: func(f *fields) {
				f.time = time.Now().Add(defaultTokenTTL).Unix()
			},
			args: args{
				userID: 1,
			},
			assert: func(t *testing.T, f *fields, service *service.AuthService, tokenStr string, expiresAt int64, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, expiresAt, f.time)

				// this part is needed to verify a valid token
				userID, err := service.ParseToken(tokenStr)
				require.NoError(t, err)
				require.Equal(t, 1, userID)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var fields fields
			test.setup(&fields)

			authService := service.NewAuthService(nil, nil)
			tokenStr, expiresAt, err := authService.GenerateToken(test.args.userID)

			test.assert(t, &fields, authService, tokenStr, expiresAt, err)
		})
	}
}

func TestAuthService_ParseToken(t *testing.T) {
	validAccessToken, _, _ := service.NewAuthService(nil, nil).GenerateToken(1)

	type args struct {
		accessToken string
	}
	tests := []struct {
		name   string
		args   args
		assert func(*testing.T, int, error)
	}{
		{
			name: "success",
			args: args{
				accessToken: validAccessToken,
			},
			assert: func(t *testing.T, userID int, err error) {
				require.NoError(t, err)
				require.Equal(t, 1, userID)
			},
		},
		{
			name: "fail - empty token",
			args: args{
				accessToken: "",
			},
			assert: func(t *testing.T, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "token contains an invalid number of segments")
				require.Equal(t, 0, userID)
			},
		},
		{
			name: "fail - signature is invalid",
			args: args{
				accessToken: validAccessToken + "test",
			},
			assert: func(t *testing.T, userID int, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "signature is invalid")
				require.Equal(t, 0, userID)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			authService := service.NewAuthService(nil, nil)
			userID, err := authService.ParseToken(test.args.accessToken)

			test.assert(t, userID, err)
		})
	}
}
