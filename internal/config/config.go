package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_HOST"`
	DBPassword string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_HOST"`
	ServerPort string `mapstructure:"DB_HOST"`
}

var (
	once sync.Once
	cfg  *Config
)

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			panic("Ошибка чтения конфига")
		}

		cfg = &Config{}
		if err := viper.Unmarshal(cfg); err != nil {
			panic("Ошибка парсинга конфига")
		}
	})

	return cfg
}
