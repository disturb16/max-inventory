package main

import (
	"context"

	"github.com/disturb/max-inventory/database"
	"github.com/disturb/max-inventory/internal/repository"
	"github.com/disturb/max-inventory/internal/service"
	"github.com/disturb/max-inventory/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),

		fx.Invoke(),
	)

	app.Run()
}
