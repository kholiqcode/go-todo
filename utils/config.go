package utils

import (
	"time"

	"github.com/spf13/viper"
)

type BaseConfig struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerPort           string        `mapstructure:"SERVER_PORT"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	SGKey                string        `mapstructure:"SG_KEY"`
	Environment          string        `mapstructure:"ENVIRONMENT"`
	ServiceName          string        `mapstructure:"SERVICE_NAME"`
	REDIS_HOST           string        `mapstructure:"REDIS_HOST"`
	REDIS_USERNAME       string        `mapstructure:"REDIS_USERNAME"`
	REDIS_PASSWORD       string        `mapstructure:"REDIS_PASSWORD"`
	CACHE_DURATION       time.Duration `mapstructure:"CACHE_DURATION"`
	MIGRATION_URL        string        `mapstructure:"MIGRATION_URL"`
}

func LoadBaseConfig(path string, configName string) (config *BaseConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	LogAndPanicIfError(err, "failed when reading config")

	err = viper.Unmarshal(&config)
	LogAndPanicIfError(err, "failed when unmarshal config")

	return
}

func CheckAndSetConfig(path string, configName string) *BaseConfig {
	config := LoadBaseConfig(path, configName)
	if config.Environment == TEST {
		config = LoadBaseConfig(path, "test")
	}

	return config
}
