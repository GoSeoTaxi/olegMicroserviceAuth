package grpcService

import (
	"context"
	"time"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerService) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	queryBuilder := sq.Select(
		"id",
		"user_name",
		"email",
		"role_id",
		"created_at",
		"updated_at").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	res := &desc.GetResponse{
		UserInfo: &desc.UserInfo{},
	}

	var createdAt, updatedAt time.Time
	err = s.DBPool.
		QueryRow(ctx, sqlQuery, args...).
		Scan(&res.UserInfo.Id,
			&res.UserInfo.Name,
			&res.UserInfo.Email,
			&res.UserInfo.Role,
			&createdAt,
			&updatedAt)
	if err != nil {
		return nil, err
	}
	res.UserInfo.CreatedAt = timestamppb.New(createdAt)
	res.UserInfo.UpdateAt = timestamppb.New(updatedAt)

	return res, nil
}
