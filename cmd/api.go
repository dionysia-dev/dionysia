package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/learn-video/streaming-platform/internal/api"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Run API server",
		Run: func(cmd *cobra.Command, args []string) {
			app := fx.New(
				fx.Provide(api.NewServer),
				fx.Invoke(func(e *gin.Engine) {}),
			)

			app.Run()
		},
	}
}
