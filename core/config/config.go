package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

/*
	All config should be required.
	Optional only allowed if zero value of the type is expected being the default value.
	time.Duration units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”. as in time.ParseDuration().
*/

type (
	Postgres struct {
		UserName string `mapstructure:"PG_USER_NAME" validate:"required"`
		Password string `mapstructure:"PG_PASSWORD" validate:"required"`
		Host     string `mapstructure:"PG_HOST" validate:"required"`
		Port     string `mapstructure:"PG_PORT" validate:"required"`
		Database string `mapstructure:"PG_DATABASE" validate:"required"`
	}

	Swagger struct {
		SwaggerHost string `mapstructure:"SWAGGER_HOST" validate:"required"`
	}

	Configuration struct {
		Environment   string `mapstructure:"ENV" validate:"required,oneof=development staging production"`
		BindAddress   int    `mapstructure:"BIND_ADDRESS" validate:"required"`
		TokenLifeSpan int    `mapstructure:"TOKEN_HOUR_LIFESPAN" validate:"required"`
		APISecret     string `mapstructure:"API_SECRET" validate:"required"`

		ServiceName string   `mapstructure:"SERVICE_NAME"`
		ServiceHost string   `mapstructure:"SERVICE_HOST"`
		Postgres    Postgres `mapstructure:",squash"`
		Swagger     Swagger  `mapstructure:",squash"`
	}
)

func InitConfig() (*Configuration, error) {
	var cfg Configuration

	viper.SetConfigType("env")
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	_, err := os.Stat(envFile)
	if !os.IsNotExist(err) {
		viper.SetConfigFile(envFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
