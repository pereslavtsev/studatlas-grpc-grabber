package config

import (
	"context"

	log "github.com/sirupsen/logrus"
	v "github.com/spf13/viper"
)

type contextKeyType string

var contextKey = contextKeyType("config")

// Config is the config operations wrapper
type Config struct {
	MongoURI string
}

func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)
	// set defaults
	v.SetDefault("mongo_uri", "mongodb://localhost:27020/test")

	v.SetConfigName(".env")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		log.Warnf("Error reading a development config file, %s, default values has been used instead.", err)
	}

	return context.WithValue(ctx, contextKey, &Config{
		MongoURI: v.GetString("mongo_uri"),
	})
}

// FromContext returns the database wrapper from a given context.
func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(contextKey).(*Config)
	if !ok {
		log.Panicf("calling config.FromContext from a non-database context")
	}
	return cfg
}
