package main

import (
	"github.com/2pizzzza/cryptosync/internal/app"
	"github.com/2pizzzza/cryptosync/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	app.New(cfg)
}
