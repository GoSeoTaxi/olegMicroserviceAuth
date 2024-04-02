package auth

import (
	"context"
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	if id < 1 {
		return &model.User{}, fmt.Errorf("no user ID information")
	}
	var user *model.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = s.authRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		logWrite := fmt.Sprintf("Get ID=%v", id)

		errTx = s.authRepository.WriterLog(ctx, logWrite)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
