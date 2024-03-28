package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	serviceMocks "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
)

func TestUpdate(t *testing.T) {
	//	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = int64(gofakeit.Uint8())
		idZero = int64(-1)
		name   = gofakeit.Name()
		email  = gofakeit.Email()

		serviceErr = fmt.Errorf("ID must be greater than 0")

		req = &desc.UpdateRequest{
			Id:    id,
			Name:  &wrapperspb.StringValue{Value: name},
			Email: &wrapperspb.StringValue{Value: email},
			Role:  1,
		}
		reqZero = &desc.UpdateRequest{
			Id:    idZero,
			Name:  &wrapperspb.StringValue{Value: name},
			Email: &wrapperspb.StringValue{Value: email},
			Role:  1,
		}

		user = &model.User{
			ID: id,
			UserT: struct {
				Name         string
				Email        string
				HashPassword string
				Role         int64
			}{Name: name, Email: email, Role: 1},
		}

		userZero = &model.User{
			ID: idZero,
			UserT: struct {
				Name         string
				Email        string
				HashPassword string
				Role         int64
			}{Name: name, Email: email, Role: 1},
		}
	)

	defer mc.Finish()

	tests := []struct {
		name            string
		args            args
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, user).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: reqZero,
			},
			err: serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, userZero).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//	t.Parallel()
			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			_, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
