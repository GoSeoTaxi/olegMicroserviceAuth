package tests

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
	serviceMocks "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/mocks"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = int64(gofakeit.Uint8())
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, false, false, 8)

		hash           = md5.Sum([]byte(password))
		hashedPassword = hex.EncodeToString(hash[:])

		serviceErr = fmt.Errorf("error create User")

		req = &desc.CreateRequest{
			UserCreate: &desc.UserInfoCreate{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
				Role:            1,
			},
		}

		info = &model.User{
			UserT: struct {
				Name         string
				Email        string
				HashPassword string
				Role         int64
			}{Name: name, Email: email, HashPassword: hashedPassword, Role: 1},
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
