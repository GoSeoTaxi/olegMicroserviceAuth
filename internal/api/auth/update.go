package auth

import (
	"context"
	"fmt"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	ex "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth/converter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	// Проверяем, установлено ли поле Name
	var hasName bool
	if req.Name != nil {
		hasName = len(req.Name.Value) > 0
	}

	// Проверяем, установлено ли поле Email
	var hasEmail bool
	if req.Email != nil {
		hasEmail = len(req.Email.Value) > 0
	}

	// Проверяем, установлено ли поле Role
	var hasRole bool
	if req.Role != 0 {
		hasRole = true
	}

	// Если нет изменений для обновления
	if !hasName && !hasEmail && !hasRole {
		return nil, fmt.Errorf("no changes provided for updating")
	}

	err := i.authService.Update(ctx, ex.ConvertProtoToModelUpdateRequest(req))

	return &emptypb.Empty{}, err
}
