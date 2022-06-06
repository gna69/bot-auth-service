package grpc_server

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

var cfg struct {
	port uint8 `env:"PORT" envDefault:"8080"`
}

func init() {
	if err := env.Parse(&cfg); err != nil {
		log.Err(err)
	}
}
