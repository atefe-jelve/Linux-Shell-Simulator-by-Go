package databases

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "postgres"
	dbUser = "tecmint"
	dbPass = "securep@wd"
)

var db *gorm.DB

func init() {
	fmt.Println("========================================================================================================================")
	fmt.Println("connecting to database...")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPass)

	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Warn,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Enable color
		},
	)

	sesssion, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		err = fmt.Errorf("unable to connect to database: %w", err)
		panic(err)
	}

	db = sesssion

	// fmt.Printf("connected to database: %s\n", dsn)
	fmt.Println("========================================================================================================================")
}

func GetDB() *gorm.DB {
	return db
}

// func SetDB(database *gorm.DB) {
// 	db = database
// }
