package auth

import (
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/client/db"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/repository"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(authRepository repository.AuthRepository, txManager db.TxManager) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
