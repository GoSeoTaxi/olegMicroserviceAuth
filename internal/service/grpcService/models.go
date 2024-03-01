package grpcService

import (
	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type serverService struct {
	desc.UnimplementedUserV1Server
	DBPool *pgxpool.Pool
}

func newServerService(pool *pgxpool.Pool) desc.UserV1Server {
	return &serverService{DBPool: pool}
}

func RegisterWithServer(s *grpc.Server, pool *pgxpool.Pool) {
	reflection.Register(s)
	desc.RegisterUserV1Server(s, newServerService(pool))
}
