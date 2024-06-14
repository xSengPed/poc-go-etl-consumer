package configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

type Configs struct {
	KafkaConfig    KafkaConfig
	PostgresConfig PostgresConfig
}

type KafkaConfig struct {
	Host string `env:"KAFKA_HOST"`
	Port string `env:"KAFKA_PORT"`
}

type PostgresConfig struct {
	Host     string `env:"PG_HOST"`
	Username string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	DBName   string `env:"PG_DBNAME"`
}

func NewConfigs() *Configs {
	err := godotenv.Load("./.env")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	configs := new(Configs)
	err = env.Parse(configs)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	slog.Info("Kafka Host: " + configs.KafkaConfig.Host)
	slog.Info("Kafka Port: " + configs.KafkaConfig.Port)
	slog.Info("Postgres Host: " + configs.PostgresConfig.Host)
	slog.Info("Postgres Username: " + configs.PostgresConfig.Username)

	return configs
}
