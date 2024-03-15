package repository

import (
	"context"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
)

type AuthRepository interface {
	Create(context.Context, *model.User) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Update(context.Context, *model.User) error
	Delete(context.Context, int64) error
	WriterLog(context.Context, string) error
}
