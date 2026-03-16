package configs

import (
	"strings"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.pass", "postgres")
	viper.SetDefault("database.database", "app")
}

func Load() error {

	// Permite usar DATABASE_HOST etc
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.BindEnv("database.host", "POSTGRES_HOST")
	viper.BindEnv("database.port", "POSTGRES_PORT")
	viper.BindEnv("database.user", "POSTGRES_USER")
	viper.BindEnv("database.pass", "POSTGRES_PASSWORD")
	viper.BindEnv("database.database", "POSTGRES_DB")

	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Pass:     viper.GetString("database.pass"),
			Database: viper.GetString("database.database"),
		},
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
