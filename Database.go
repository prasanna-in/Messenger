package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type Database interface {
	createBot(Mbot) string
	lastBotCreated() string
	Allbots() string
}

type DB struct {
	*gorm.DB

}

func Create_Db_Connection(connectionString string) *DB {
	db,err := gorm.Open("postgres",connectionString)
	if err != nil{
		log.Fatal("Database Connection did not happen ....")
	}
	return &DB{db}

}
//type Mbot struct {
//	gorm.Model
//	Name         string
//	Sendid       int
//	Secretstring string
//}


//type Dbcon *gorm.DB
//
//var Db Dbcon
//
// Global method not going to use this so discarding it ...
//func Get_Database_Connection() *gorm.DB {
//	if Db == nil {
//		db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
//		if err != nil {
//			log.Fatal(err)
//		}
//		Db = db
//		return Db
//	}else {return Db}
//}
//func(m Mbot) Create() string {
//	db := Get_Database_Connection()
//	var mbot1 Mbot
//	db.Where("Secretstring = ?", m.Secretstring).First(&mbot1)
//	if mbot1.Secretstring == m.Secretstring {
//		log.Println("User secret alredy used try another one ....")
//		return "Token already registered request for a new one from shifu@thoughtworks.com"
//
//	}else {
//		db.NewRecord(m)
//		db.Create(&m)
//		log.Println("Group Created ....")
//		return "Group has been created... "
//	}
//
//	return "1"
//}