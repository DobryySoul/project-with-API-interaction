package main

import (
	"DobryySoul/project-with-API-interaction/internal/app"
	"context"
)

func main() {
	ctx := context.Background()

	app.Run(ctx)
}
