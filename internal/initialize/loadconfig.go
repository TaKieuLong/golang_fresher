package initialize

import (
	"fmt"

	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/spf13/viper"
)

func Loadconfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to docode configuration %v", err)
	}
}