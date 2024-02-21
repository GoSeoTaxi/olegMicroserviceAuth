package grpcService

import (
	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ServerService struct {
	desc.UnimplementedUserV1Server
	DBPool *pgxpool.Pool
}
