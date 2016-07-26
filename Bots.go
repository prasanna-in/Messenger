package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Mbot struct {
	gorm.Model
	Name         string
	Sendid       int
	Secretstring string
}

func (db *DB) createBot(m Mbot) string {
	var mbot1 Mbot
	db.Where("Secretstring = ?", m.Secretstring).First(&mbot1)
	if mbot1.Secretstring == m.Secretstring {
		log.Println("This Group has already been created....")
		return "This Secret has been taken consider using a different secret ID ... "
	}else {
		db.NewRecord(m)
		db.Create(&m)
		log.Println("Group Created ....")
		return "Group has been created... "
	}
	return "This shoud never hit ...."

}

func (db *DB) lastBotCreated() string  {
	var mbot1 Mbot
	db.Last(&mbot1)
	return "The Last Group created is "+mbot1.Name+" with Secretkey "+mbot1.Secretstring
}

func (db *DB) Allbots() string {
	rows,_ := db.Rows()
	log.Println(rows)
	return  "    "
}