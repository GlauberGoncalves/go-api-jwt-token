package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {

	banco := os.Getenv("DB_DRIVER")

	if banco == "sqlite3" {
		db, err := gorm.Open("sqlite3", "/banco.db")
		if err != nil {
			fmt.Errorf("erro ao conectar com database")
		}
		return db
	}

	if banco == "mysql" {
		db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
		if err != nil {
			fmt.Errorf("erro ao conectar com database")
		}
		return db
	}

	return nil
}
