package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	user := os.Getenv("MYSQL_DB_USER")
	password := os.Getenv("MYSQL_DB_PASSWORD")
	host := os.Getenv("MYSQL_DB_HOST")
	dbName := os.Getenv("MYSQL_DB")

	connection, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbName+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}

	db = connection
}

func GetMysqlConnectionPool() *gorm.DB { return db }
