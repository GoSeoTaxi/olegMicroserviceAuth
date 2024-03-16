package user

import (
	"context"
	"fmt"
	"time"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Update(ctx context.Context, user *model.User) error {

	queryBuilder := sq.Update("users").PlaceholderFormat(sq.Dollar)
	if len(user.UserT.Name) > 0 {
		queryBuilder = queryBuilder.Set("user_name", user.UserT.Name)
	}
	if len(user.UserT.Email) > 0 {
		queryBuilder = queryBuilder.Set("email", user.UserT.Email)
	}
	if user.UserT.Role > 0 {
		queryBuilder = queryBuilder.Set("role_id", user.UserT.Role)
	}

	queryBuilder = queryBuilder.Set("updated_at", time.Now()).Where(sq.Eq{idColumn: user.ID})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "note_repository.Update",
		QueryRaw: sqlQuery,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}
