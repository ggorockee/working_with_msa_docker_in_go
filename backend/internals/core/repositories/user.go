package repositories

import (
	"back-end/internals/core/ports"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

var _ ports.UserRepository = (*UserRepository)(nil)

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

func getEnv(key string, defaultValue string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defaultValue
	}

	return envValue

}

func NewUserRepository() *UserRepository {

	dbConfig := DBConfig{
		getEnv("MYSQL_USER", "ggorockee"),
		getEnv("MYSQL_PASSWORD", "ggorockee"),
		getEnv("MYSQL_HOST", "localhost"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_DATABASE", "ggorockee"),
	}

	db, err := gorm.Open(mysql.Open(dbConfig.String()), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect to database \n", err)
		os.Exit(2)
	}

	log.Println("successfully db connected!")

	if err := db.AutoMigrate(
	// models
	); err != nil {
		log.Println("failed to migrate to database\n", err)
	}

	return &UserRepository{
		conn: db,
	}
}

func (r *UserRepository) Register(email string, password string) error {
	requestPayload := struct {
		email    string
		password string
	}{
		email:    email,
		password: password,
	}
	if err := r.conn.Create(&requestPayload).Error; err != nil {
		return err
	}

	return nil
}
