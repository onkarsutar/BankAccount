package confighelper

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("./configs/config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file %s ", err))
	}
}

func GetConfig(keyName string) string {
	keyValue := viper.GetString(keyName)
	return keyValue
}
