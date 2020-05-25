package config

import (
	"context"

	log "github.com/sirupsen/logrus"
	v "github.com/spf13/viper"
)

var contextKey = "config"

// Config is the config operations wrapper
type Config struct {
	Address  string
	MongoURI string
}

func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)
	// set defaults
	v.SetDefault("address", ":50051")
	v.SetDefault("mongo_uri", "mongodb://localhost:27020/test")

	v.SetConfigName(".env")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		log.Warnf("Error reading a development config file, %s, default values has been used instead.", err)
	}

	return context.WithValue(ctx, contextKey, &Config{
		Address:  v.GetString("address"),
		MongoURI: v.GetString("mongo_uri"),
	})
}

// FromContext returns the database wrapper from a given context.
func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(contextKey).(*Config)
	if !ok {
		log.Fatal("calling config.FromContext from a non-database context")
	}
	return cfg
}
