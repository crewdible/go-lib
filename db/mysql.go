package db

import (
	"errors"
	"fmt"
	"os"

	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init - mysql init
func InitMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err = gorm.Open(mysql.Open(
		dsn, // data source name
	), &gorm.Config{})

	if err != nil {
		return errors.New("MySQL Connection Error")
	}

	return err
}

// DbManager - return db connection
func MySQLManager() *gorm.DB {
	return db
}
