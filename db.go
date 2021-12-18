package main

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// Declare the table for the database
type BlacklistedIP struct {
	gorm.Model
	IP   string `gorm:"type:varchar(16);uniqueIndex"`
	Path string `gorm:"type:varchar(2048)"`
}

// ConnectDB connect to db
func ConnectDB() error {
	var err error
	p := GetEnv("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Cannot parse POSTGRES_PORT. Will use default port: 5432")
		port = 5432
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("POSTGRES_HOST"), port, GetEnv("POSTGRES_USER"),
		GetEnv("POSTGRES_PASSWORD"), GetEnv("POSTGRES_DB"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Println("Connected to database")
	}

	// Create or migrate the schema
	if DB.AutoMigrate(&BlacklistedIP{}); err != nil {
		log.Fatal(err)
	} else {
		log.Default().Println("Schema successfully created or migrated")
	}
	return nil
}

// Check if the IP is blacklisted
func IsBlacklisted(ip string) bool {
	var blacklistedIP BlacklistedIP

	DB.Where("ip = ?", ip).First(&blacklistedIP)
	return blacklistedIP.IP != ""
}

// Put IP, Path and DateTime to the database
func BlacklistIP(ip string, path string) {
	var blacklistedIP BlacklistedIP
	blacklistedIP.IP = ip
	blacklistedIP.Path = path
	if !IsBlacklisted(ip) {
		err := DB.Create(&blacklistedIP).Error
		if err != nil {
			log.Fatal(err)
		}
		log.Default().Println("IP: " + ip + " blacklisted")
		SendMail(ip)
	}
}
