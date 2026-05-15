package main

import (
	"context"
	"log"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err = a.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
