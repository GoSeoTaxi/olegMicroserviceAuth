package auth

import (
	"context"
	"fmt"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	ex "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/api/auth/converter"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userInfo := req.GetUserCreate()
	if req.UserCreate == nil {
		return nil, fmt.Errorf("no user information provided")
	}

	if userInfo.Name == "" ||
		userInfo.Email == "" ||
		userInfo.Password == "" ||
		userInfo.PasswordConfirm == "" ||
		userInfo.Role < 1 {
		return nil, fmt.Errorf("not all required fields provided")
	}

	if userInfo.Password != userInfo.PasswordConfirm {
		return nil, fmt.Errorf("passwords are not equal")
	}

	user := ex.ConvertProtoToModelCreateRequest(req)

	if user.ID != 0 {
		return nil, fmt.Errorf("request has ID")
	}

	res, err := i.authService.Create(ctx, user)
	result, err := ex.ConvertModelToProtoCreateRequest(res)
	if err != nil {
		return nil, err
	}
	return result, err
}
