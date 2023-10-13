package database

import (
	"fmt"
	"glamour_reserve/app/config"
	"glamour_reserve/entity/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDBMysql(cfg *config.DBConfig) *gorm.DB {
	
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASSWORD, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitMigrate(DB)

	return DB
}

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{}, &models.Booking{}, &models.DetailBooking{}, &models.Service{}, models.Variant{})
}
