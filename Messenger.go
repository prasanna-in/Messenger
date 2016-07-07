package main

import (
	"net/http"
	"os"
	"log"
)

func HttpHandler2(response http.ResponseWriter, request *http.Request)  {
	log.Println("Token Call Received")
}
func main() {
	http.HandleFunc("/testing123", HttpHandler2)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
