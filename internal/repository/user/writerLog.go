package user

import (
	"context"
	"log"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/client/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) WriterLog(ctx context.Context, logStr string) error {

	queryBuilder := sq.Insert("log_table").
		PlaceholderFormat(sq.Dollar).
		Columns("action").
		Values(logStr)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "log.Writer",
		QueryRaw: sqlQuery,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("err write to log: %v\n", err)
		return err
	}

	return nil
}
