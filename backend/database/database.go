package database

import (
	"back-end/configs"
	"back-end/internals/core/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBConn struct {
	Conn *gorm.DB
}

var DB DBConn

func Connect() *DBConn {
	config := configs.New()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect to database \n", err)
		os.Exit(2)
	}

	log.Println("successfully db connected!")
	if err := db.AutoMigrate(
		// models
		new(domain.User),
		new(domain.Memo),
	); err != nil {
		log.Println("failed to migrate to database\n", err)
	}

	DB = DBConn{Conn: db}
	return &DB
}
