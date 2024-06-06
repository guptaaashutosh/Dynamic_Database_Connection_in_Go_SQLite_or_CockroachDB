package config

import (
	"fmt"
	"log"
	"path"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Port          string `mapstructure:"PORT"`
	DB_CONNECTION string `mapstructure:"DB_CONNECTION"`
	COCKROACH_DB_URL string `mapstructure:"COCKROACH_DB_URL"`
}

func Load() (Config, error) {

	gp := "C:/Users/ZTI/go"
	ap := path.Join(gp, "/src/tutorials/dynamic_db_conn_project/.config")

	fmt.Println(".config path: ", ap)

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(ap)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("no config file found. going on...")
		} else {
			log.Printf("error loading config file")
		}
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return Config{}, err
	}

	return Conf, err
}
