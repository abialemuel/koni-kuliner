package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/subosito/gotenv"
)

var create_table_products = "CREATE TABLE IF NOT EXISTS products (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`name` VARCHAR(55) NOT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

func main() {
	_ = gotenv.Load()
	port := os.Getenv("DATABASE_PORT")
	if port == "" {
		port = "3306"
	}

	dataSourceName := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(port)) + ")/" + os.Getenv("DATABASE_NAME")
	log.Println("CONNECTING TO:", os.Getenv("DATABASE_HOST")+":"+(string(port)))
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		defer db.Close()
	}

	log.Println("MIGRATING...")
	if err != nil {
		panic(err.Error())
	}

	db.MustExec(create_table_products)
	log.Println("DONE")
}
