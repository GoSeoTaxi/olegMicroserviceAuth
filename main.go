package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/cmd"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.HiMagentaString("Preparing..."))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go cmd.RunService(ctx)

	<-shutdown
	fmt.Println(color.HiRedString("Shutting down..."))

	cancel()
}
