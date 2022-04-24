package db

import (
	util "demo/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=demo password=postgres sslmode=disable")
	if err != nil {
		util.PrintError(err)
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func GetTable(table string) *gorm.DB {
	return db.Table(table)
}
