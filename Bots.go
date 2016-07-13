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

func createBot(db *gorm.DB,m *Mbot) string {
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