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
	"strings")

//type Env struct {
//	db *gorm.DB
//
//}

type Env struct {
	db Database

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
	var validID = regexp.MustCompile(`Register \d\d\d\d\d\d`)
	if validID.MatchString(text) == true {
		Actualtext := strings.Split(text, " ")[1]
		db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		var mbot1 Mbot
		db.Where("Secretstring = ?", Actualtext).First(&mbot1)
		if mbot1.Secretstring == Actualtext {
			if mbot1.Sendid == 0 {
				log.Println("The Group ID is : " + strconv.Itoa(Telegramresponse.Message.Chat.Id))
				mbot1.Sendid = Telegramresponse.Message.Chat.Id
				db.Save(&mbot1)
				SendmessageInternal(mbot1.Sendid, "Group Registered, You can add more users and have fun.")
			}else {
				log.Println("Token Already Registered ....")
				SendmessageInternal(Telegramresponse.Message.Chat.Id, "Token already registered request new from shifu@thoughtworks.com")
			}
		}
	}else {
		log.Println("We dont care about anyother message ....")
	}

}

func (e *Env) Dbcreate(response http.ResponseWriter, request *http.Request) {
	Name1 := request.FormValue("Name")
	Secret := request.FormValue("secret")
	mb := Mbot{Name:Name1,Secretstring:Secret}
	str := e.db.createBot(mb)
	fmt.Fprintln(response,str)
}


func (e *Env) Dbview(response http.ResponseWriter, request *http.Request) {
	str := e.db.lastBotCreated()
	fmt.Fprintf(response,str)

}

func messageSender(sendID int, message string)  {
	str := fmt.Sprintf("https://api.telegram.org/"+os.Getenv("Bot_API")+"/sendMessage?chat_id=%d&text=%s&parse_mode=Markdown", sendID, message)
	log.Println(str)
	resp, err := http.Get(str)
	log.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func Sendmessage(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Server", "GO_Messenger_Bot")
	response.WriteHeader(200)
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
		messageSender(mbot1.Sendid,Message)
	}
}
func (e Env) All(response http.ResponseWriter,request *http.Request)  {
	str := e.db.Allbots()
	for _,bot := range(str) {
		fmt.Fprintf(response,"%s, %s, %d\n\n", bot.Name,bot.Secretstring,bot.Sendid)

	}
}
func main() {
	db := Create_Db_Connection(os.Getenv("DATABASE_URL"))
	env := &Env{db}
	http.HandleFunc("/testing123", TelegramHandler)
	http.HandleFunc("/Create", env.Dbcreate)
	http.HandleFunc("/view", env.Dbview)
	http.HandleFunc("/All", env.All)
	http.HandleFunc("/Sendmessage",Sendmessage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
