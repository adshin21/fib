package main

import (
	"fmt"

	"github.com/adshin21/fib/config"
	"github.com/adshin21/fib/server"
)

func main() {
	fmt.Println("Hello")

	config.Init()
	server.Init()

	fmt.Printf("%+v\n", config.Cfg)
}
