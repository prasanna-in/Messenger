//Author : Prasanna Kanagasabai

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
	"regexp"
	"strings"
)

type Mbot struct {
	gorm.Model
	Name         string
	Sendid       int
	Secretstring string
}

func SendmessageInternal(sendid int, str string) {
	str1 := fmt.Sprintf("https://api.telegram.org/"+os.Getenv("Bot_API")+"/sendMessage?chat_id=%d&text=%s&parse_mode=Markdown", sendid, str)
	http.Get(str1)
}

func TelegramHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Telegram Handler ...")
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
	var Mbot1 Mbot
	db.Where("Secretstring = ?",text).First(&Mbot1)
	log.Println(Mbot1.Secretstring+"JK")
	var validID = regexp.MustCompile(`Register \d\d\d\d\d\d`)
	log.Println(validID.MatchString(text))
	if validID.MatchString(text) == true {
		mbo2 := Mbot1
		Actualtext := strings.Split("register 123456", " ")[1]
		log.Println(Actualtext)
		log.Println(mbo2.Secretstring+"pK")
		if Mbot1.Secretstring == Actualtext {
			if Mbot1.Sendid == 0 {
				log.Println("The Group ID is : " + strconv.Itoa(Telegramresponse.Message.Chat.Id))
				Mbot1.Sendid = Telegramresponse.Message.Chat.Id
				db.Save(&Mbot1)
				SendmessageInternal(Mbot1.Sendid, "Group Registered, You can add more users and have fun.")
			}else {
				log.Println("Token Already Registered ....")
				SendmessageInternal(Telegramresponse.Message.Chat.Id, "Token already registered request new from shifu@thoughtworks.com")
			}
		}
	}else {
		log.Println("Nothing I am bothered with ..... Heee :)")
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
		fmt.Fprintf(response,"Token already registered request for a new one from shifu@thoughtworks.com")

	}else {
	log.Println(secretstring)
	mbot := Mbot{Name: name, Secretstring: secretstring}
	db.NewRecord(mbot)
	db.Create(&mbot)
	//log.Println(db.NewRecord(mbot))
	log.Println("Group Created ....")
	fmt.Fprintf(response,"Group has been created... ")
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

		str := fmt.Sprintf("https://api.telegram.org/"+os.Getenv("Bot_API")+"/sendMessage?chat_id=%d&text=%s&parse_mode=Markdown", mbot1.Sendid, Message)
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

func main() {
	http.HandleFunc("/testing123", TelegramHandler)
	http.HandleFunc("/Create", Dbcreate)
	http.HandleFunc("/view", Dbview)
	http.HandleFunc("/Sendmessage",Sendmessage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
