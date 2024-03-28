package tests

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
	serviceMocks "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = int64(gofakeit.Uint8())

		name  = gofakeit.Name()
		email = gofakeit.Email()

		password       = gofakeit.Password(true, true, true, false, false, 8)
		hash           = md5.Sum([]byte(password))
		hashedPassword = hex.EncodeToString(hash[:])

		req = &desc.GetRequest{
			Id: id,
		}
		res = &desc.GetResponse{
			UserInfo: &desc.UserInfo{
				Id:    id,
				Name:  name,
				Email: email,
				Role:  1,
				CreatedAt: &timestamppb.Timestamp{
					Seconds: 111,
					Nanos:   2,
				},
				UpdateAt: &timestamppb.Timestamp{
					Seconds: 222,
					Nanos:   1,
				},
			},
		}

		info = &model.User{
			ID: id,
			UserT: struct {
				Name         string
				Email        string
				HashPassword string
				Role         int64
			}{Name: name, Email: email, HashPassword: hashedPassword, Role: 1},
			CreatedAt: time.Unix(res.UserInfo.CreatedAt.Seconds, int64(res.UserInfo.CreatedAt.Nanos)),
			UpdatedAt: time.Unix(res.UserInfo.UpdateAt.Seconds, int64(res.UserInfo.UpdateAt.Nanos)),
		}

		idZero     = int64(-1)
		serviceErr = fmt.Errorf("no valid id")

		req1 = &desc.GetRequest{
			Id: idZero,
		}

		userZero = &model.User{}
		resZero  = &desc.GetResponse{UserInfo: nil}
	)

	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
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
				mock.GetMock.Expect(ctx, id).Return(info, nil)
				return mock
			},
		},

		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: req1,
			},
			want: resZero,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, idZero).Return(userZero, serviceErr)
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
			responseGet, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, responseGet)

		})
	}
}
