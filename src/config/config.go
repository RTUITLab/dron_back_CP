package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DB		*DBConfig
	Auth	*Auth
}

type DBConfig struct {
	URI 	string 	`envconfig:"CP_BACK_DBURI" default:"root:root@/site"`
}

type Auth struct {
	AccessSecret	string	`envconfig:"CP_BACK_ACCESS_SECRET" default:"access_secret"`
	RefreshSecret	string	`envconfig:"CP_BACK_REFRESH_SECRET" default:"refresh_secret"`
}

func GetConfig() *Config {
	var config Config 
	{
		if err := godotenv.Load("./.env"); err != nil {
			log.Warn("Don't find env file")
		}

		if err := envconfig.Process("cp_back", &config); err != nil {
			log.WithFields(
				log.Fields{
					"function" : "envconfig.Process",
					"error"	:	err,
				},
			).Fatal("Can't read env vars, shutting down...")
		}
	}

	return &config
}