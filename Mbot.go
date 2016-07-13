package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Mbot struct {
	gorm.Model
	Name         string
	Sendid       int
	Secretstring string
}

type Dbcon struct {
	con *gorm.DB
}
var Db Dbcon

func db() *gorm.DB {
	if Db.con == nil {
		db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}
		Db.con = db
		return Db.con
	}else {return Db.con}
}
func(m Mbot) Create() string {
	db := db()
	var mbot1 Mbot
	db.Where("Secretstring = ?", m.Secretstring).First(&mbot1)
	if mbot1.Secretstring == m.Secretstring {
		log.Println("User secret alredy used try another one ....")
		return "Token already registered request for a new one from shifu@thoughtworks.com"

	}else {
		db.NewRecord(m)
		db.Create(&m)
		log.Println("Group Created ....")
		return "Group has been created... "
	}

	return "1"
}