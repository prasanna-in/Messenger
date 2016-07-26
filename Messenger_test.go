package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

type mockDB struct {}

func (mdb *mockDB) Allbots()[]Mbot  {
	bots := make([]Mbot, 0)
	bots = append(bots,Mbot{Secretstring:"Test",Sendid:45,Name:"PK"})
	//bots = append(bots,Mbot{Secretstring:"Joke",Sendid:23,Name:"Joke"})
	return bots
}

func (mdb *mockDB)createBot(Mbot)string  {
	return " "

}
func (mdb *mockDB)lastBotCreated()string  {
	return " "

}

func TestAll(test *testing.T)  {
	rec := httptest.NewRecorder()
	req,_ := http.NewRequest("GET","/All",nil)
	env := Env{db:&mockDB{}}
	http.HandlerFunc(env.All).ServeHTTP(rec, req)
	expected := "PK, Test, 45\n\n"
	if expected != rec.Body.String(){
		test.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
	
}

//func TestDbview(test *testing.T)  {
//	rec := httptest.NewRecorder()
//	req,_ := http.NewRequest("GET","/View",nil)
//
//
//}