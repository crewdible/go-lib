package db

import (
	"errors"
	"fmt"

	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
)

var db = make(map[string]*gorm.DB)
var err error

// Init - mysql init
func InitMySQL(connName, dbUser, dbPassword, dbHost, dbPort, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	db[connName], err = gorm.Open(mysql.Open(
		dsn, // data source name
	), &gorm.Config{})

	if err != nil {
		msg := fmt.Sprintf("Connection %s => MySQL Connection Error", connName)
		return errors.New(msg)
	}

	return err
}

// DbManager - return db connection
func MySQLManager() map[string]*gorm.DB {
	return db
}
