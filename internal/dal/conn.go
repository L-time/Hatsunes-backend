package dal

import (
	"HatsuBlog/internal/configs"
	"HatsuBlog/pkg/log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConn() *gorm.DB {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn,
		configs.GlobalConfig.Host,
		configs.GlobalConfig.User,
		configs.GlobalConfig.Password,
		configs.GlobalConfig.Database,
		configs.GlobalConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log.NewOrmLogger(),
	})
	if err != nil {
		panic(err)
	}
	return db
}
