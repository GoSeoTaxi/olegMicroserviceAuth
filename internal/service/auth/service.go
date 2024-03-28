package auth

import (
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/repository"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service"
	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db"
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

func NewMockService(userRepo repository.AuthRepository, txManager db.TxManager) service.AuthService {
	return &serv{
		authRepository: userRepo,
		txManager:      txManager,
	}
}
