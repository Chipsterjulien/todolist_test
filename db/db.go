package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Messages struct {
	Id    int    `gorm:"primary_key"`
	Title string `sql:"title"`
	//Id    int `sql:"unique"`
	//Title string
}

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "todolist.db")
	if err != nil {
		log.Fatal("Unable to open todolist.db", err)
	}
	db.LogMode(true)

	db.CreateTable(new(Messages))

	return &db
}
