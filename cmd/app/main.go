// Package main is the entrypoint to ws
package main

import (
	"github.com/adshin21/fib/config"
	"github.com/adshin21/fib/internal/app"
)

func main() {
	cfg := config.MustGet()
	app.Run(cfg)
}
