package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/config"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service/grpcService"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunService start service Auth
func RunService(ctx context.Context) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("parcing config err = ", err)
	}

	log.Println(color.HiYellowString("Init DB..."))
	dbDSN := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=%v", cfg.HostDB, cfg.PortDB, cfg.NameDB, cfg.LoginDB, cfg.PasswordDB, cfg.SSLTypeDB)
	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	log.Println(color.HiRedString("Starting..."))

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.HostDB, cfg.PortGRPC))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &grpcService.ServerService{DBPool: pool})

	log.Printf(color.HiYellowString("server listening at %v \n", lis.Addr()))

	go func() {
		<-ctx.Done()
		gracefulStopServer(s)
	}()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
