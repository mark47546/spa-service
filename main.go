package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

)

//Main Function
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to our Spa Service")
	fmt.Println("homePage")
}

func handleRequest(){
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/sessions", returnAllSessions)
	log.Fatal(http.ListenAndServe(":10000", r))

}

func main(){
	fmt.Println("Rest API v2.0 - Mux Routers")
	SpaSessions = []SpaSession{
		SpaSession{Id:"1",Session:"Herbal Spa", Time:"8:00-10:00",Description:"A"},
		SpaSession{Id:"2",Session:"Thai Massage", Time:"10:00-12:00",Description:"B"},
	}
	handleRequest()
}

//Additional Function
func returnAllSessions(w http.ResponseWriter, r *http.Request){
	fmt.Println("returnAllSessions")
	json.NewEncoder(w).Encode(SpaSessions)
}



type SpaSession struct {
	Id string `json:"Id"`
	Session string `json:"Session"`
	Time string `json:"Time"`
	Description string `json:"Description"`
}

var SpaSessions []SpaSession

