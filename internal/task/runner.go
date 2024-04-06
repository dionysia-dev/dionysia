package task

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/learn-video/streaming-platform/internal/config"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, cfg *config.Config, srv *asynq.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			mux := asynq.NewServeMux()
			mux.HandleFunc(TypeStreamPackage, HandleStreamPackageTask)

			go func() {
				if err := srv.Run(mux); err != nil {
					log.Fatalf("Could not start server: %s\n", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown()
			return nil
		},
	})
}
