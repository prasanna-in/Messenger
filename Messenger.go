package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"strconv"
)

type Mbot struct {
	gorm.Model
	Name         string
	Sendid       int
	Secretstring string
}

func HttpHandler2(response http.ResponseWriter, request *http.Request) {
	log.Println("Token Call Received")
	var Telegramresponse TGUpdate
	bodystring := request.Body
	body, err := ioutil.ReadAll(bodystring)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &Telegramresponse)
	text := Telegramresponse.Message.Text
	log.Println(text)
	switch text {
	case "/Register":
		var val int
		//log.Println("The Group ID is : " + strconv.Itoa(Telegramresponse.Message.Chat.Id))
		val = Telegramresponse.Message.Chat.Id
		log.Println("The Group ID is : " + strconv.Itoa(val))
	default:
		log.Println(Telegramresponse.Message.From.Username)
	}
}
func HttpHandler(response http.ResponseWriter, request *http.Request) {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)
	db.AutoMigrate(&Mbot{})
	//db.Create(Mbot{Name: "PK", Sendid: 123, Secretstring: "sdkjaskdjh"})
	//var mbot Mbot
	//log.Println(db.First(&mbot, 1))

}
func main() {
	http.HandleFunc("/testing123", HttpHandler2)
	http.HandleFunc("/Create", HttpHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
