package main

import (
	"github.com/roka-crew/samsamoohooh-backend/internal/postgres"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/handler"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/middleware"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/token"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Supply("./configs/config.yaml"),
		fx.Provide(
			config.New,
			postgres.New,

			store.NewUserStore,
			store.NewGroupStore,
			store.NewGoalStore,
			store.NewTopicStore,

			service.NewUserSerivce,
			service.NewGroupService,
			service.NewGoalService,
			service.NewTopicService,
			service.NewAuthService,

			token.NewJWTMaker,
			middleware.NewAuthMiddleware,
			server.NewServer,
		),
		fx.Invoke(
			handler.NewUserHandler,
			handler.NewGroupHandler,
			handler.NewGoalHander,
			handler.NewTopicHandler,
			handler.NewAuthHandler,
		),
	).Run()
}
