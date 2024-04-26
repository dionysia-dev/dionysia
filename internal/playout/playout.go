package playout

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func New(cfg *config.Config) *gin.Engine {
	e := gin.Default()

	e.Static("/playout", cfg.PlayoutPath)

	return e
}

func registerHooks(lc fx.Lifecycle, cfg *config.Config, e *gin.Engine) {
	srv := &http.Server{
		Addr:              ":" + cfg.PlayoutPort,
		Handler:           e,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Printf("Failed starting server at: %s", srv.Addr)
				return err
			}

			go srv.Serve(ln) //nolint:errcheck // error is handled by shutdown

			fmt.Printf("Server started at: %s", srv.Addr)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				log.Printf("Failed to shutdown server: %v", err)
				return err
			}

			fmt.Println("Server stopped")

			return nil
		},
	})
}
