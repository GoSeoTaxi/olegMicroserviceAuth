package user

import (
	"context"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	queryBuilder := sq.Insert("users").PlaceholderFormat(sq.Dollar).
		Columns("user_name", "email", "password_hash", "role_id").
		Values(user.UserT.Name, user.UserT.Email, user.UserT.HashPassword, user.UserT.Role).
		Suffix("RETURNING id")

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: sqlQuery,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}
