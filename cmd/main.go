package main

import (
	"github.com/osg_task/internal/app"
	"github.com/osg_task/internal/pkg/config"
)

func main() {
	cfg := config.Load()
	app.Run(&cfg)
}