package auth

import (
	"context"
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
)

func (s *serv) Update(ctx context.Context, req *model.User) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.authRepository.Update(ctx, req)
		if errTx != nil {
			return errTx
		}

		logWrite := fmt.Sprintf("Update req=%+v", req)
		errTx = s.authRepository.WriterLog(ctx, logWrite)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
