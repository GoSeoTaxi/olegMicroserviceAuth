package auth

import (
	"context"
	"fmt"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	ex "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth/converter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	if req.Id < 1 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	userID := ex.ConvertProtoToModelDeleteRequest(req)
	err := i.authService.Delete(ctx, userID)
	return &emptypb.Empty{}, err
}
