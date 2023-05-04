package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	dirCurrent, _ := os.Getwd()
	dirRoot := filepath.Join(dirCurrent, "./")
	log.Print("====", dirRoot)
	log.Print("====", dirCurrent)

	viper.SetConfigFile("/.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Unable to decode into struct", err)
	}

	return &env
}
