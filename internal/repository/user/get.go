package user

import (
	"context"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/client/db"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	queryBuilder := sq.Select(
		idColumn,
		titleColumn,
		mailColumn,
		hashColumn,
		roleColumn,
		createdAtColumn,
		updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	res := &model.User{}
	q := db.Query{
		Name:     "auth_repository.Get",
		QueryRaw: sqlQuery,
	}
	err = r.db.DB().
		QueryRowContext(ctx, q, args...).
		Scan(&res.ID,
			&res.UserT.Name,
			&res.UserT.Email,
			&res.UserT.HashPassword,
			&res.UserT.Role,
			&res.CreatedAt,
			&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return res, nil
}
