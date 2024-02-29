package grpcService

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	sq "github.com/Masterminds/squirrel"
)

func (s *serverService) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userInfo := req.GetUserCreate()
	if userInfo == nil {
		return nil, fmt.Errorf("no user information provided")
	}

	if userInfo.Name == "" ||
		userInfo.Email == "" ||
		userInfo.Password == "" ||
		userInfo.PasswordConfirm == "" ||
		userInfo.Role < 1 {
		return nil, fmt.Errorf("not all required fields provided")
	}

	if userInfo.Password != userInfo.PasswordConfirm {
		return nil, fmt.Errorf("passwords are not equal")
	}

	hash := md5.Sum([]byte(userInfo.GetPassword()))
	hashedPassword := hex.EncodeToString(hash[:])

	queryBuilder := sq.Insert("users").PlaceholderFormat(sq.Dollar).
		Columns("user_name", "email", "password_hash",
			"role_id").
		Values(userInfo.GetName(), userInfo.GetEmail(), hashedPassword,
			int64(userInfo.GetRole().Number())).
		Suffix("RETURNING id")

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var id int64
	err = s.DBPool.QueryRow(ctx, sqlQuery, args...).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
