package cmd

import (
	"log"

	"github.com/dionysia-dev/dionysia/internal/api"
	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/dionysia-dev/dionysia/internal/db"
	"github.com/dionysia-dev/dionysia/internal/db/redistore"
	"github.com/dionysia-dev/dionysia/internal/logging"
	"github.com/dionysia-dev/dionysia/internal/queue"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Run API server",
		Run: func(*cobra.Command, []string) {
			if err := godotenv.Load(".env"); err != nil {
				log.Println("Could not load env file")
			}

			app := fx.New(
				fx.Provide(
					config.New,
					logging.New,
					db.New,
					db.NewDBInputStore,
					queue.NewClient,
					redistore.NewClient,
					fx.Annotate(
						redistore.NewOriginStore,
						fx.As(new(service.OriginStore)),
					),
					service.NewOriginHandler,
					service.NewNotificationHandler,
				),
				api.Module,
			)

			app.Run()
		},
	}
}
