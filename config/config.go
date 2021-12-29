package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerGrpc GrpcConfig
	ServerHttp HttpConfig
	Redis RedisConfig
}

type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

type GrpcConfig struct {
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
}

type HttpConfig struct {
	Port string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
}

func LoadConfig(configName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(configName)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, err
}

func GetConfigPath(configPath string) string {
	if configPath == "some" {
		return "some path"
	}

	return "./config/config"
}
