package service

import (
	"context"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, info *model.User) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Update(context.Context, *model.User) error
	Delete(context.Context, int64) error
}
