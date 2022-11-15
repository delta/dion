package repository

import (
	"fmt"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// connects to postgres db and runs migrations
func Init() {
	fmt.Println("==>INITIALIZING DATABASE")
	if err := connect(); err != nil {
		errMsg := fmt.Errorf("unable to connect to db because %+v", err)
		panic(errMsg)
	}
	fmt.Println("==>RUNNING MIGRATIONS")
	if err := autoMigrate(); err != nil {
		errMsg := fmt.Errorf("unable to run migrations due to %+v", err)
		panic(errMsg)
	}
}

// connects to the database using config, and returns error if any
func connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		config.C.Db.Host,
		config.C.Db.User,
		config.C.Db.Password,
		config.C.Db.DbName,
		config.C.Db.Port,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = conn
	return err
}

// runs auto-migration, and returns error if any
func autoMigrate() error {
	err := db.AutoMigrate(&models.User{})
	return err
}
