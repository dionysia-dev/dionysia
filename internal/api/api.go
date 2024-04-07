package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn-video/streaming-platform/internal/config"
	"github.com/learn-video/streaming-platform/internal/db"
	"github.com/learn-video/streaming-platform/internal/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(dbq *db.Queries, _ *zap.SugaredLogger, nh service.NotificationHandler) *gin.Engine {
	inputController := NewInputController(service.NewInputHandler(dbq))
	notificationController := NewNotificationController(nh)

	e := gin.Default()
	e.POST("/inputs", inputController.CreateInput)
	e.GET("/inputs/:id", inputController.GetInput)
	e.DELETE("/inputs/:id", inputController.DeleteInput)
	e.POST("/notifications/package", notificationController.EnqueuePackaging)

	return e
}

func registerHooks(lc fx.Lifecycle, cfg *config.Config, e *gin.Engine) {
	srv := &http.Server{
		Addr:              ":" + cfg.APIPort,
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
