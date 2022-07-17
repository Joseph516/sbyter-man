package repository

import (
	"douyin_service/services/user/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		config.DatabaseCfg.UserName,
		config.DatabaseCfg.Password,
		config.DatabaseCfg.Host,
		config.DatabaseCfg.DBName,
		config.DatabaseCfg.Charset,
		config.DatabaseCfg.ParseTime,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
