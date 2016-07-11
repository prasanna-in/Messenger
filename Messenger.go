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
	"fmt"
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
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var mbot1 Mbot
	db.Where("Secretstring = ?",text).First(&mbot1)
	if mbot1.Secretstring == text{
		if mbot1.Sendid == 0 {
			log.Println("The Group ID is : " + strconv.Itoa(Telegramresponse.Message.Chat.Id))
			mbot1.Sendid = Telegramresponse.Message.Chat.Id
			db.Save(&mbot1)
		}else {
			log.Println("Token Already Registered ....")
		}
	}

}

func Dbcreate(response http.ResponseWriter, request *http.Request) {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name := request.FormValue("Name")
	secretstring := request.FormValue("secret")
	var mbot1 Mbot
	db.Where("Secretstring = ?", secretstring).First(&mbot1)
	if mbot1.Secretstring == secretstring {
		log.Println("User secret alredy used try another one ....")

	}else {
	log.Println(secretstring)
	mbot := Mbot{Name: name, Secretstring: secretstring}
	db.NewRecord(mbot)
	db.Create(&mbot)
	log.Println(db.NewRecord(mbot))
}

}
func Dbview(response http.ResponseWriter, request *http.Request) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var mbot1 Mbot
	db.Last(&mbot1)
	log.Println(mbot1.Name,mbot1.Secretstring,mbot1.Sendid)

}
func Sendmessage(response http.ResponseWriter, request *http.Request) {
	Secrestring := request.FormValue("secret")
	Message := request.FormValue("text")
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var mbot1 Mbot
	db.Where("Secretstring = ?",Secrestring).First(&mbot1)
	log.Println(mbot1)
	if mbot1.Secretstring == Secrestring{
		str := fmt.Sprintf("https://api.telegram.org/bot249456369:AAHaHfsSSkiiEPeiwqnChNX16sbS4H-JHqM/sendMessage?chat_id=%d&text=%s&parse_mode=Markdown", mbot1.Sendid, Message)
		log.Println(str)
		resp, err := http.Get(str)
		log.Println(resp)
		if err != nil {
			log.Fatal(err)
		}
	}
	response.Header().Set("Server", "GO_Messenger_Bot")
	response.WriteHeader(200)
}
func Probe(response http.ResponseWriter, request * http.Request)  {

	str := "https://api.telegram.org/bot249456369:AAHaHfsSSkiiEPeiwqnChNX16sbS4H-JHqM/sendMessage?chat_id=-122886380&text=Check"
	resp,_:= http.Get(str)
	log.Println(resp)

}
func main() {
	http.HandleFunc("/testing123", HttpHandler2)
	http.HandleFunc("/Create", Dbcreate)
	http.HandleFunc("/view", Dbview)
	http.HandleFunc("/Sendmessage",Sendmessage)
	http.HandleFunc("/probe" , Probe)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
