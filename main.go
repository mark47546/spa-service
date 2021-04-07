package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	r.HandleFunc("/session",createSession).Methods("POST")
	r.HandleFunc("/session/{id}", deleteSession).Methods("DELETE")
	r.HandleFunc("/session/{id}", returnSingleSession)
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

func createSession(w http.ResponseWriter, r *http.Request){
    reqBody, _ := ioutil.ReadAll(r.Body)
    var session SpaSession
    json.Unmarshal(reqBody, &session)
    SpaSessions = append(SpaSessions, session)
    json.NewEncoder(w).Encode(session)
}

func returnSingleSession(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    for _, session := range SpaSessions {
        if session.Id == key{
            json.NewEncoder(w).Encode(session)
        }
    }
}

func deleteSession(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    for index, session := range SpaSessions{
        if session.Id == id{
            SpaSessions = append(SpaSessions[:index],SpaSessions[index+1:]...)
        }
    }
}



type SpaSession struct {
	Id string `json:"Id"`
	Session string `json:"Session"`
	Time string `json:"Time"`
	Description string `json:"Description"`
}

var SpaSessions []SpaSession

