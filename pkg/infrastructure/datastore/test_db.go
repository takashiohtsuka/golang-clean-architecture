package datastore

import (
	"log"

	mysqlcfg "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewTestDB() *gorm.DB {
	mySqlConfig := &mysqlcfg.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "golang_clean_architecture_test",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := gorm.Open(mysql.Open(mySqlConfig.FormatDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
