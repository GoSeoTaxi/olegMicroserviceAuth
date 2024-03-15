package auth

import (
	"context"
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
)

func (s *serv) Create(ctx context.Context, req *model.User) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.authRepository.Create(ctx, req)
		if errTx != nil {
			return errTx
		}

		logWrite := fmt.Sprintf("Create user=%+v", req)

		errTx = s.authRepository.WriterLog(ctx, logWrite)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, err
}
