package main

import "fmt"
import (
	"net/http"
	"os"
)

func HttpHandler2(response http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(response,"Welcome")
	
}
func main() {
	http.HandleFunc("/testing123", HttpHandler2)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
