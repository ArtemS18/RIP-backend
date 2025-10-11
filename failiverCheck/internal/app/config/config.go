package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MinioConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
}
type JWTConfig struct {
	SecretKey        string `mapstructure:"secret_key"`
	ExpiresAtMinutes int    `mapstructure:"expire_at_minutes"`
}

type Config struct {
	Minio  *MinioConfig  `mapstructure:"minio"`
	Server *ServerConfig `mapstructure:"server"`
	JWT    *JWTConfig    `mapstructure:"jwt"`
}

func NewConfig() (*Config, error) {
	configName := "config"
	_ = godotenv.Load()
	if envName := os.Getenv("CONFIG_NAME"); envName != "" {
		configName = envName
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	serverConf := &ServerConfig{}
	minioConf := &MinioConfig{}
	config := &Config{Server: serverConf, Minio: minioConf}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	// err = viper.Unmarshal(minioConf)
	// if err != nil {
	// 	return nil, err
	// }
	log.Info(config.Minio.AccessKey)
	log.Info("Config created")
	return config, nil
}
