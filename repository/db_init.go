package repository

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {

	var err error

	dsn := "root:rJukuUgN@tcp(172.16.32.4:49362)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err

}
