package grpcService

import (
	"context"
	"fmt"
	"time"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerService) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	if req.Id < 1 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	if !(len(req.Name.Value) > 0 || len(req.Email.Value) > 0 || req.GetRole().Number() > 0) {
		return nil, fmt.Errorf("no changes provided for updating")
	}

	queryBuilder := sq.Update("users").PlaceholderFormat(sq.Dollar)
	if len(req.Name.Value) > 0 {
		queryBuilder = queryBuilder.Set("user_name", req.Name.Value)
	}
	if len(req.Email.Value) > 0 {
		queryBuilder = queryBuilder.Set("email", req.Email.Value)
	}
	if req.GetRole().Number() > 0 {
		queryBuilder = queryBuilder.Set("role_id", int64(req.Role.Number()))
	}

	queryBuilder = queryBuilder.Set("updated_at", time.Now()).Where(sq.Eq{"id": req.GetId()})

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
		return nil, fmt.Errorf("no rows were updated")
	}

	return &emptypb.Empty{}, nil
}
