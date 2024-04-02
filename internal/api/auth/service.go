package auth

import (
	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
