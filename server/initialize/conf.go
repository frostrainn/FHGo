package initialize

import (
	"fhgo/config"
	"fhgo/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	c := config.Conf{}
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")
	err := v.Unmarshal(&c)
	if err != nil {
		return
	}
	global.Config = &c

	//viper.SetConfigName("config.yaml")
}
