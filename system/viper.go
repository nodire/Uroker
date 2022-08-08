package system

import (
	"github.com/spf13/viper"
	"log"
)

func GetConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
