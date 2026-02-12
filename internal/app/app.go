// Package app
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/adshin21/fib/config"
	"github.com/adshin21/fib/pkg/httpserver"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func Run(cfg *config.Config) {
	handler := gin.New()
	setupRoutes(handler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	server := httpserver.New(handler, httpserver.Port(cfg.Server.Port))

	g.Go(func() error {
		select {
		case err := <-server.Notify():
			return fmt.Errorf("app - Run - server.Notify: %w", err)
		case <-ctx.Done():
			return nil // Context was cancelled by signal, this is a clean exit
		}
	})

	g.Go(func() error {
		<-ctx.Done() // Wait for the interrupt signal
		return server.Shutdown()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("app - Run - errgroup.Wait: %v\n", err)
	}
}
