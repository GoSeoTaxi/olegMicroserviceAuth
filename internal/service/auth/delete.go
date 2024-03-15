package auth

import (
	"context"
	"fmt"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.authRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		logWrite := fmt.Sprintf("Delete ID=%v", id)

		errTx = s.authRepository.WriterLog(ctx, logWrite)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
