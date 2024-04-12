package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/learn-video/dionysia/internal/config"
	"github.com/learn-video/dionysia/internal/logging"
	"github.com/learn-video/dionysia/internal/queue"
	"github.com/learn-video/dionysia/internal/task"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewWorker() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Run worker server",
		Run: func(*cobra.Command, []string) {
			if err := godotenv.Load(".env"); err != nil {
				log.Println("Could not load env file")
			}

			app := fx.New(
				fx.Provide(config.New),
				fx.Provide(logging.New),
				fx.Provide(queue.NewServer),
				fx.Invoke(task.Run),
			)

			app.Run()
		},
	}
}
