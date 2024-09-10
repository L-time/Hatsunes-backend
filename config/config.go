package config

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed config.yaml
var configFile []byte

func ReleaseConfigFile() {
	file, err := os.OpenFile("config.yaml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic("创建配置文件失败！请检查文件系统是否正常！")
	}
	defer file.Close()
	_, err = file.Write(configFile)
	if err != nil {
		panic("配置文件写入失败！请检查文件是否成功创建！")
	}
	fmt.Println("配置文件生成成功！请修改后重新启动程序。")
	os.Exit(0)
}
