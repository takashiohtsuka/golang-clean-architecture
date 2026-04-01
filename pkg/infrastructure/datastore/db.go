package datastore

import (
	"log"

	"golang-clean-architecture/pkg/config"

	mysqlcfg "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
*
mysqlの設定とgormの設定
*/
func NewDB() *gorm.DB {
	mySqlConfig := &mysqlcfg.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(mysql.Open(mySqlConfig.FormatDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
