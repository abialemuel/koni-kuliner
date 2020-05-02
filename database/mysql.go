package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

func InitMysql() *gorm.DB {
	_ = gotenv.Load()
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	fmt.Println("initializing mysql connection")

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully initialize mysql connection")

	return db
}

// func InitMysql() (*sqlx.DB, error) {
// 	_ = gotenv.Load()
// 	dbPort := os.Getenv("DATABASE_PORT")
// 	if dbPort == "" {
// 		dbPort = "3306"
// 	}

// 	dataSourceName := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(dbPort)) + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"
// 	db, err := sqlx.Open("mysql", dataSourceName)

// 	if err != nil {
// 		return nil, err
// 	}

// 	db.SetConnMaxLifetime(time.Minute * 10)

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Successfully initialize mysql connection")

// 	return db, err
// }
