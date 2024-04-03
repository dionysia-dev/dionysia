package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn-video/streaming-platform/internal/config"
	"github.com/learn-video/streaming-platform/internal/db"
	"github.com/learn-video/streaming-platform/internal/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(dbq *db.Queries, logger *zap.SugaredLogger) *gin.Engine {
	inputController := NewInputController(service.NewInputHandler(dbq))

	e := gin.Default()
	e.POST("/inputs", inputController.CreateInput)

	return e
}

func registerHooks(lc fx.Lifecycle, cfg *config.Config, e *gin.Engine) {
	srv := &http.Server{Addr: ":8080", Handler: e}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Printf("Failed starting server at: %s", srv.Addr)
				return err
			}

			go srv.Serve(ln)

			fmt.Printf("Server started at: %s", srv.Addr)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			fmt.Println("Server stopped")
			return nil
		},
	})
}
