package cmd

import (
	"log"

	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/dionysia-dev/dionysia/internal/playout"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewPlayoutServer() *cobra.Command {
	return &cobra.Command{
		Use:   "playout",
		Short: "Run a playout server",
		Run: func(*cobra.Command, []string) {
			if err := godotenv.Load(".env"); err != nil {
				log.Println("Could not load env file")
			}

			app := fx.New(
				fx.Provide(config.New),
				playout.Module,
			)

			app.Run()
		},
	}
}
