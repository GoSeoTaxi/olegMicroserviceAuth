package auth

import (
	"context"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	ex "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth/converter"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userID := ex.ConvertProtoToModelGetRequest(req)
	res, err := i.authService.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := ex.ConvertModelToProtoGetRequest(res)
	return result, nil
}
