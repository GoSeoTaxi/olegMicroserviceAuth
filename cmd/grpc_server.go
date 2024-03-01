package cmd

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

const timeOutBeforeClose = 5 * time.Second

func gracefulStopServer(s *grpc.Server) {
	log.Println("Stopping gRPC server...")
	stopped := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		log.Println("gRPC server stopped gracefully")
	case <-time.After(timeOutBeforeClose):
		log.Println("gRPC server stopping timeout reached")
		s.Stop()
	}
}
