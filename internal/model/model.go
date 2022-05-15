package model

import (
	"context"
	"douyin_service/pkg/setting"
	"fmt"
	"github.com/mattn/go-colorable"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var newLogger = logger.New(
		log.New(colorable.NewColorableStdout(), "", log.LstdFlags),
		logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		return nil, err
	}

	//db.Callback().Create().Before("gorm:create").Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Before("gorm:update").Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Before("gorm:delete").Replace("gorm:delete", deleteCallback)
	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	mysqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	nowTime := time.Now().Unix()
	if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
		_ = createTimeField.Set(context.Background(), db.Statement.ReflectValue, nowTime)
	}

	if modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn"); modifyTimeField != nil {
		_ = modifyTimeField.Set(context.Background(), db.Statement.ReflectValue, nowTime)
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	field := "ModifiedOn"
	if timeField := db.Statement.Schema.LookUpField(field); timeField != nil {
		_ = timeField.Set(context.Background(), db.Statement.ReflectValue, time.Now().Unix())
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
