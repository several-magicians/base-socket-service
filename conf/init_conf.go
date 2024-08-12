package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Port string = ""

	Net string = ""
)

func InitConfig() {
	viper.SetConfigFile(GetYaml())
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Port = viper.GetString("port")

	Net = viper.GetString("net")

}
