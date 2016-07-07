package main

import (
	"net/http"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

func HttpHandler2(response http.ResponseWriter, request *http.Request)  {
	log.Println("Token Call Received")
	var Telegramresponse TGUpdate
	bodystring := request.Body
	body, err := ioutil.ReadAll(bodystring)
	if err != nil{
		log.Fatal(err)
	}
	json.Unmarshal(body,&Telegramresponse)
	text := Telegramresponse.Message.Text
	switch text {
	case "/Register":
		//log.Fatal(Telegramresponse.Message.Chat.ID())
	default:
		log.Fatal("Default Called")
	}
}
func main() {
	http.HandleFunc("/testing123", HttpHandler2)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
