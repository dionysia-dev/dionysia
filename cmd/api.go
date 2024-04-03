package cmd

import (
	"github.com/learn-video/streaming-platform/internal/api"
	"github.com/learn-video/streaming-platform/internal/config"
	"github.com/learn-video/streaming-platform/internal/db"
	"github.com/learn-video/streaming-platform/internal/logging"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Run API server",
		Run: func(cmd *cobra.Command, args []string) {
			app := fx.New(
				fx.Provide(config.New),
				fx.Provide(logging.New),
				fx.Provide(db.NewPool),
				fx.Provide(db.NewQuerier),
				api.Module,
			)

			app.Run()
		},
	}
}
