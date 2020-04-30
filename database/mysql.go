package database

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // for sqlx
	"github.com/jmoiron/sqlx"
	"github.com/subosito/gotenv"
)

func InitMysql() (*sqlx.DB, error) {
	_ = gotenv.Load()
	dbPort := os.Getenv("DATABASE_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dataSourceName := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(dbPort)) + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"
	db, err := sqlx.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully initialize mysql connection")

	return db, err
}
