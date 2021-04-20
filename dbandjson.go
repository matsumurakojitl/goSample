package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"log"
	//"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbinit()
	dbInsert("todo", "未達成")
}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func dbinit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Print(err)
		fmt.Print(err)
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Print(err)
		fmt.Print(err)
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}
