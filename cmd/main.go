package main

import (
	"context"
	"flag"
	"log"
	"test-task/internal/app"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "config", "local.env", "Path to config")
	flag.Parse()
	ctx := context.Background()
	a, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = a.RunHttp()
	log.Fatal(err)
}
