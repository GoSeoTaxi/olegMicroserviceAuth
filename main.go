package main

import (
	"context"
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/cmd"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.HiMagentaString("Preparing..."))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd.RunService(ctx)

	fmt.Println(color.HiRedString("Shutting down..."))

}
