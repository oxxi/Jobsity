package database

import (
	"fmt"
	"log"
	"os"

	"github.com/oxxi/jobsity/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")              // your database user
	password := getEnv("DB_PASSWORD", "123456789") // your database password
	database := getEnv("DB_NAME", "jobsity")       // name of you database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	d.AutoMigrate(&models.Task{})
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
