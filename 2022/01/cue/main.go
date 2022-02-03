package main

import (
	"fmt"

	"github.com/luthermonson/tutorials-2022-01-cue/config"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("sample")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("unable to read the config file")
	}

	var c config.Config
	err = viper.Unmarshal(&c)
	if err != nil {
		panic("unable to decode into struct")
	}

	fmt.Printf("%+v", c)
}
