package cmd

import (
	"context"
	"log"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/app"
)

// RunService start service Auth
func RunService(ctx context.Context) {

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

}
