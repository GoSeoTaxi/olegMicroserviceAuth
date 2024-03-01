package grpcService

import (
	"context"
	"fmt"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serverService) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req.Id < 1 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	queryBuilder := sq.Delete("users").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": req.Id})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := s.DBPool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows were deleted")
	}

	return &emptypb.Empty{}, nil
}
