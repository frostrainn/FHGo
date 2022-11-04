package initialize

import "github.com/spf13/viper"

func InitConfig() {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigType("yaml")
	v.SetConfigName("test.yaml")

	//viper.SetConfigName("config.yaml")
}
