package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // for sqlx
	"github.com/jinzhu/gorm"
)

func InitMysql() *gorm.DB {
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
	defer db.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully initialize mysql connection")

	return db
}
