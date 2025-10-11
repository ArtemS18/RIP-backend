package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
	_ = godotenv.Load()
	prefix := "APP"
	keySeparator := "__"

	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", keySeparator))
	viper.AutomaticEnv()

	configName := os.Getenv("CONFIG_NAME")
	if configName == "" {
		configName = "config"
	}
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); !notFound {
			return nil, err
		}
	}
	LoadEnvInViper(prefix, keySeparator)

	cfg := &Config{
		Server: &ServerConfig{},
		Minio:  &MinioConfig{},
		JWT:    &JWTConfig{},
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if cfg.JWT.SecretKey == "" {
		return nil, fmt.Errorf("jwt.secret_key is required (ENV or file)")
	}
	return cfg, nil
}

func LoadEnvInViper(prefix string, keySeparator string) {
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		}
		name, val := parts[0], parts[1]
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		key := strings.TrimPrefix(name, prefix)
		key = strings.TrimPrefix(key, keySeparator)
		key = strings.ReplaceAll(key, keySeparator, ".")
		key = strings.ToLower(key)
		viper.Set(key, val)
	}
}
