package config

import (
	"fmt"
	"log"
	"path"
	"tutorials/dynamic_db_conn_project/constants"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Port          string `mapstructure:"PORT"`
	DB_TYPE string `mapstructure:"DB_TYPE"`
	COCKROACH_DB_URL string `mapstructure:"COCKROACH_DB_URL"`
}


// Define a custom "DbType" type to represent the enum
type DbType string

const (
  DbTypeSqlite  DbType = constants.SQLITE_DB
  DbTypeCockroach DbType = constants.COCKROACH_DB
)

// Validation function to check if the value matches a valid DbType
func isValidDbType(value string) bool {
	for _, dt := range []DbType{DbTypeSqlite, DbTypeCockroach} {
	  if value == string(dt) {
		return true
	  }
	}
	return false
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

	// Validate the DB_TYPE value
	if !isValidDbType(string(Conf.DB_TYPE)) {
		return Config{}, fmt.Errorf("invalid DB_TYPE value: %s", Conf.DB_TYPE)
	}

	return Conf, err
}
