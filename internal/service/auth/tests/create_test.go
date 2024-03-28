package tests

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/auth"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/mocks"
	// "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/auth"

	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db"
	dbMocks "github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db/mocks"
	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db/pg"
	"github.com/GoSeoTaxi/olegMicroservicePlatform/pkg/db/transaction"

	repoMocks "github.com/GoSeoTaxi/olegMicroserviceAuth/internal/repository/mocks"
	// converterLogRepo "github.com/ybgr111/auth/internal/repository/log/converter"
	// converterUserRepo "github.com/ybgr111/auth/internal/repository/user/converter"
)

type createUserVariables struct {
	userInfo *model.User
}

type CreateUserSuite struct {
	ctx       context.Context
	ctxWithTx context.Context

	suite.Suite

	mc                 *minimock.Controller
	userRepositoryMock *repoMocks.AuthRepositoryMock
	//	logRepositoryMock  *repoMocks.LogRepositoryMock
	fakeTxMock     *mocks.FakeTxMock
	transactorMock *dbMocks.TransactorMock

	txManagerMock db.TxManager

	createUserVariables
}

func TestCreateUserSuite(t *testing.T) {
	suite.Run(t, new(CreateUserSuite))
}

func (s *CreateUserSuite) SetupSuite() {
	s.ctx = context.Background()
	s.mc = minimock.NewController(s.T())

	s.userRepositoryMock = repoMocks.NewAuthRepositoryMock(s.mc)
	// s.logRepositoryMock = repoMocks.NewLogRepositoryMock(s.mc)
	s.fakeTxMock = mocks.NewFakeTxMock(s.mc)

	s.ctxWithTx = pg.MakeContextTx(s.ctx, s.fakeTxMock)

	s.transactorMock = dbMocks.NewTransactorMock(s.mc)
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	s.transactorMock.BeginTxMock.Expect(s.ctx, txOpts).Return(s.fakeTxMock, nil)
	s.txManagerMock = transaction.NewTransactionManager(s.transactorMock)

	s.userInfo = &model.User{}

	s.userInfo.ID = int64(gofakeit.Uint8())
	name := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, false, false, 8)

	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	s.userInfo = &model.User{
		UserT: struct {
			Name         string
			Email        string
			HashPassword string
			Role         int64
		}{Name: name, Email: email, HashPassword: hashedPassword, Role: 1},
	}

}

func (s *CreateUserSuite) TestCreate_Success() {
	// Специфичные моки методов.
	s.userRepositoryMock.CreateMock.Expect(s.ctxWithTx, s.userInfo).Return(s.userInfo.ID, nil)
	logWrite := fmt.Sprintf("Create user=%+v", s.userInfo)
	s.userRepositoryMock.WriterLogMock.Expect(s.ctxWithTx, logWrite).Return(nil)
	s.fakeTxMock.CommitMock.Return(nil)

	service := auth.NewMockService(s.userRepositoryMock, s.txManagerMock)

	newID, err := service.Create(s.ctx, s.userInfo)

	// Проверки корректности теста.
	require.Equal(s.T(), nil, err)
	require.Equal(s.T(), s.userInfo.ID, newID)
}
