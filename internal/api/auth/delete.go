package auth

import (
	"context"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	ex "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth/converter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	userID := ex.ConvertProtoToModelDeleteRequest(req)
	err := i.authService.Delete(ctx, userID)
	return &emptypb.Empty{}, err
}
