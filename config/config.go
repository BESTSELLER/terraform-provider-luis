package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

// EnvConfig defines the structure of the global configuration parameters
type EnvConfig struct {
	Username string `required:"false"`
	Password string `required:"false"`
	LogLevel string `required:"false"`
	Port     int    `default:"8080"`
}

// EnvVars stores the Global Configuration.
var EnvVars EnvConfig

//LoadEnvConfig Loads config from env
func LoadEnvConfig() {
	err := envconfig.Process("", &EnvVars)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read env config")
	}
}
