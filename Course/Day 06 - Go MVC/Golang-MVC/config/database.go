package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Get the correct database connection string based on the environment
func getDBConfigByEnv(env string) string {
	var user, password, host, port, name string

	switch env {
	case "dev":
		user = os.Getenv("DEV_DB_USER")
		password = os.Getenv("DEV_DB_PASSWORD")
		host = os.Getenv("DEV_DB_HOST")
		port = os.Getenv("DEV_DB_PORT")
		name = os.Getenv("DEV_DB_NAME")
	case "qc":
		user = os.Getenv("QC_DB_USER")
		password = os.Getenv("QC_DB_PASSWORD")
		host = os.Getenv("QC_DB_HOST")
		port = os.Getenv("QC_DB_PORT")
		name = os.Getenv("QC_DB_NAME")
	case "prod":
		user = os.Getenv("PROD_DB_USER")
		password = os.Getenv("PROD_DB_PASSWORD")
		host = os.Getenv("PROD_DB_HOST")
		port = os.Getenv("PROD_DB_PORT")
		name = os.Getenv("PROD_DB_NAME")
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	user, password, host, port, name)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", user, password, name, host, port)
	println(dsn)
	return dsn
}

func ConnectDB() (*gorm.DB, error) {
	env := os.Getenv("ENV")
	dsn := getDBConfigByEnv(env)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
