package tests

import (
	"context"
	"fmt"
	"testing"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
	serviceMocks "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = int64(gofakeit.Uint8())
		idZero = int64(-1)

		serviceErr = fmt.Errorf("ID must be greater than 0")

		req = &desc.DeleteRequest{
			Id: id,
		}
		reqZero = &desc.DeleteRequest{
			Id: idZero,
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
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				mock.DeleteMock.Expect(ctx, idZero).Return(serviceErr)
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
			_, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)

		})
	}

}
