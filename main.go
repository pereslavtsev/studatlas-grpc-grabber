package main

import (
	"grabber/config"
	"grabber/grabber"
	"os"

	"grabber/database"
	"grabber/server"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.SetLevel(log.DebugLevel)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func main() {
	ctx := config.Init(context.Background())
	ctx = database.Init(ctx)
	ctx = grabber.Init(ctx)
	ctx = server.Init(ctx)

	s := server.FromContext(ctx)
	s.Start()
}
