package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure: "port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:host`
	} `mapstructure"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // file name
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w \n", err))
	}

	fmt.Println("ServerPort::", viper.GetInt("server.port"))
	fmt.Println("ServerPort::", viper.GetString("security.jwt.key"))

	//config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to docode configuration %v", err)
	}

	fmt.Println("Config Port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
