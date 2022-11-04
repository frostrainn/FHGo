package initialize

import "github.com/spf13/viper"

func InitConfig() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")

	//viper.SetConfigName("config.yaml")
}
