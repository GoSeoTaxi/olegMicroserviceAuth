package auth

import (
	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
)

/*type AuthService interface {
	Create(context.Context, *model.User) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Update(context.Context, *model.User) error
	Delete(context.Context, int64) error
}*/

/*type ServiceUser struct {
	authRepo AuthService
}

func NewService(authRepo AuthService) *ServiceUser {
	return &ServiceUser{
		authRepo: authRepo,
	}
}*/

type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
