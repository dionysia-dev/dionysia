package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn-video/dionysia/docs"
	"github.com/learn-video/dionysia/internal/config"
	"github.com/learn-video/dionysia/internal/db"
	"github.com/learn-video/dionysia/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// @title           Streaming Platform API
// @version         1.0
// @description     Manage your streaming workflow
// @BasePath  /api/v1

func New(inputStore db.InputStore, _ *zap.SugaredLogger, nh service.NotificationHandler) *gin.Engine {
	inputController := NewInputController(service.NewInputHandler(inputStore))
	notificationController := NewNotificationController(nh)

	e := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := e.Group("/api/v1")
	v1.POST("/inputs", inputController.CreateInput)
	v1.GET("/inputs/:id", inputController.GetInput)
	v1.DELETE("/inputs/:id", inputController.DeleteInput)
	v1.POST("/notifications/package", notificationController.EnqueuePackaging)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
