package user

import (
	"context"
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/client/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	queryBuilder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumn: id})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "note_repository.Delete",
		QueryRaw: sqlQuery,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	return nil
}
