package configs

import (
	"HatsuBlog/config"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var GlobalConfig *DB

func checkConfigFile() bool {
	_, err := os.Stat("./config.yaml")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func LoadConfig() {
	GlobalConfig = &DB{}
	if !checkConfigFile() {
		fmt.Println("未检测到配置文件！")
		config.ReleaseConfigFile()
	}
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败！: %s \n", err))
	}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		panic(fmt.Errorf("读取配置文件失败！:%s \n", err))
	}

}
