package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

type DBConfig struct {
	userId   string
	password string
	host     string
	port     string
	dbName   string
}

func (config DBConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.userId,
		config.password,
		config.host,
		config.port,
		config.dbName,
	)
}

func ConnectDB() {
	dbConfig := DBConfig{
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	}

	db, err := gorm.Open(mysql.Open(dbConfig.String()), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("successfully db connected!")

	if err := db.AutoMigrate(
		new(User),
	); err != nil {
		log.Printf("Failed to migrate to database. \n", err)
	}

	DBConn = db
}
