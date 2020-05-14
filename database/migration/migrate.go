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
	"`brand_id` INT(11) DEFAULT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_outlets = "CREATE TABLE IF NOT EXISTS outlets (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`name` VARCHAR(55) NOT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_brands = "CREATE TABLE IF NOT EXISTS brands (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`name` VARCHAR(55) NOT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_customers = "CREATE TABLE IF NOT EXISTS customers (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`name` VARCHAR(55) NOT NULL," +
	"`address` VARCHAR(255) NOT NULL," +
	"`phone` VARCHAR(20) NOT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_outlet_products = "CREATE TABLE IF NOT EXISTS outlet_products (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`outlet_id` INT(11) DEFAULT NULL," +
	"`product_id` INT(11) DEFAULT NULL," +
	"`price` INT(11) DEFAULT NULL," +
	"`order_price` INT(11) DEFAULT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"`state` TINYINT(1) DEFAULT 0," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_cart_items = "CREATE TABLE IF NOT EXISTS cart_items (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`user_id` INT(11) DEFAULT NULL," +
	"`outlet_product_id` INT(11) DEFAULT NULL," +
	"`transaction_id` INT(11) DEFAULT NULL," +
	"`price` INT(11) DEFAULT NULL," +
	"`order_price` INT(11) DEFAULT NULL," +
	"`quantity` INT(11) DEFAULT NULL," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_transactions = "CREATE TABLE IF NOT EXISTS transactions (" +
	"`id` INT(11) NOT NULL AUTO_INCREMENT," +
	"`user_id_id` INT(11) DEFAULT NULL," +
	"`cart_item_id` INT(11) DEFAULT NULL," +
	"`amount` BIGINT DEFAULT 0," +
	"`created_at` DATETIME DEFAULT NULL," +
	"`updated_at` DATETIME DEFAULT NULL," +
	"`po_date` DATETIME DEFAULT NULL," +
	"`state` TINYINT(1) DEFAULT 0," +
	"`note` VARCHAR(255) NOT NULL," +
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
	db.MustExec(create_table_outlets)
	db.MustExec(create_table_brands)
	db.MustExec(create_table_customers)
	db.MustExec(create_table_outlet_products)
	db.MustExec(create_table_cart_items)
	db.MustExec(create_table_transactions)
	log.Println("DONE")
}
