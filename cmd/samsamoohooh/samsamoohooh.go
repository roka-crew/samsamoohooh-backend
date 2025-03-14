package main

import (
	"github.com/roka-crew/samsamoohooh-backend/internal/handler"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
	"github.com/roka-crew/samsamoohooh-backend/pkg/database/postgres"
	"github.com/roka-crew/samsamoohooh-backend/pkg/server/http"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply("./configs/config.yaml"),
		fx.Provide(
			config.New,
			postgres.New,

			store.NewUserStore,
			service.NewUserSerivce,

			http.NewServer,
		),
		fx.Invoke(
			handler.NewUserHandler,
		),
	).Run()
}
