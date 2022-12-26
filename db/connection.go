package db

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	models_tables "go_test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func Init() {
	var (
		DB_HOST    = os.Getenv("DB_HOST")
		DB_USER    = os.Getenv("DB_USER")
		DB_PASS    = os.Getenv("DB_PASS")
		DB_NAME    = os.Getenv("DB_NAME")
		DB_PORT    = os.Getenv("DB_PORT")
		DB_SSLMODE = os.Getenv("DB_SSLMODE")
		DB_TZ      = os.Getenv("DB_TZ")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT, DB_SSLMODE, DB_TZ)

	gormCustomLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1 * time.Second,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
			Colorful:                  true,
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormCustomLogger,
	})

	if err != nil {
		panic("failed to connect sql database")
	}

	clearDB := flag.Bool("clearDB", false, "a bool")

	if *clearDB {
		fmt.Println("Deleting all data from database...")
		db.Exec("DROP SCHEMA public CASCADE")
		db.Exec("CREATE SCHEMA public")
		fmt.Println("Successfully deleted all data from database...")
	}

	db.AutoMigrate(models_tables.Tables...)

	flag.Parse()

}

func DbManager() *gorm.DB {
	return db
}
