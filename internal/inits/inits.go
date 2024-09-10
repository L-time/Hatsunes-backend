package inits

import (
	"HatsuBlog/internal/configs"
	"HatsuBlog/internal/dal"
	"HatsuBlog/pkg/log"
)

func InitAll() {
	configs.LoadConfig()
	log.SysLog = log.NewSysLogger()
	log.NetLog = log.NewNetLogger()
	dal.DB = dal.NewConn()
}
