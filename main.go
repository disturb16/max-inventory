package main

import (
	"context"

	"github.com/disturb/max-inventory/database"
	"github.com/disturb/max-inventory/settings"
	"go.uber.org/fx"
)

const key string = "12345678901234567890123456789012"

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),

		fx.Invoke(),
	)

	app.Run()
}
