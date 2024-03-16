package user

import (
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/repository"
	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db"
)

const (
	tableName = "users"

	idColumn        = "id"
	titleColumn     = "user_name"
	mailColumn      = "email"
	hashColumn      = "password_hash"
	roleColumn      = "role_id"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}
