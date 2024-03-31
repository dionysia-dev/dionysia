package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle) *gin.Engine {
	e := gin.Default()

	srv := &http.Server{Addr: ":8080", Handler: e}

	addRoutes(e)

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

	return e
}
