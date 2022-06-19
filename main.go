package main

import (
	"context"
	"log"

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

		fx.Invoke(
			func(ctx context.Context, svc service.Service) {
				err := svc.RegisterUser(ctx, "joe2@email.com", "Joe", "password2")
				if err != nil {
					panic(err)
				}

				u, err := svc.LoginUser(ctx, "joe2@email.com", "password2")
				if err != nil {
					panic(err)
				}

				log.Println(u)
			},
		),
	)

	app.Run()
}
