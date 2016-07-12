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

func(m Mbot) Create(name string,secret string) string {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var mbot1 Mbot
	db.Where("Secretstring = ?", secret).First(&mbot1)
	if mbot1.Secretstring == secret {
		log.Println("User secret alredy used try another one ....")
		return "Token already registered request for a new one from shifu@thoughtworks.com"

	}else {
		mbot := Mbot{Name: name, Secretstring: secret}
		db.NewRecord(mbot)
		db.Create(&mbot)
		//log.Println(db.NewRecord(mbot))
		log.Println("Group Created ....")
		return "Group has been created... "
	}
	log.Println(m.Name+"PRasanna")
	return "1"
}