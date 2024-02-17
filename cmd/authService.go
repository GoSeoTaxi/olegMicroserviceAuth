package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/config"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	desc.UnimplementedUserV1Server
}

// RunService start service Auth
func RunService() {
	cfg := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	var role desc.Role
	switch gofakeit.Number(0, 2) {
	case 1:
		role = desc.Role_USER
	case 2:
		role = desc.Role_ADMIN
	default:
		role = desc.Role_UNKNOWN
	}

	return &desc.GetResponse{
		UserInfo: &desc.UserInfo{
			Id:        req.GetId(),
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Role:      role,
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdateAt:  timestamppb.New(gofakeit.Date()),
		},
	}, nil
}
