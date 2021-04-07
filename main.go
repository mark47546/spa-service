package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type SpaSession struct {
	Id          string `json:"Id"`
	Session     string `json:"Session"`
	Time        string `json:"Time"`
	Description string `json:"Description"`
}

var SpaSessions []SpaSession

//Main Function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our Spa Service")
	fmt.Println("homePage")
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/sessions", returnAllSessions)
	r.HandleFunc("/session", createSession).Methods("POST")
	r.HandleFunc("/session/{id}", deleteSession).Methods("DELETE")
	r.HandleFunc("/session/{id}", updateSession).Methods("PUT")
	r.HandleFunc("/session/{id}", returnSingleSession)
	log.Fatal(http.ListenAndServe(":10000", r))

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequest()
}

//Additional Function
func returnAllSessions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllSessions")
	json.NewEncoder(w).Encode(SpaSessions)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var session SpaSession
	json.Unmarshal(reqBody, &session)
	SpaSessions = append(SpaSessions, session)
	json.NewEncoder(w).Encode(session)
}

func returnSingleSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, session := range SpaSessions {
		if session.Id == key {
			json.NewEncoder(w).Encode(session)
		}
	}
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, session := range SpaSessions {
		if session.Id == id {
			SpaSessions = append(SpaSessions[:index], SpaSessions[index+1:]...)
		}
	}
}

func updateSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var tempSession SpaSession
	for index, session := range SpaSessions {
		if (session.Id) == id {
			_ = json.NewDecoder(r.Body).Decode(&tempSession)
			tempSession.Id = strconv.Itoa(index)
			SpaSessions[index] = tempSession
			json.NewEncoder(w).Encode(SpaSessions[index])
			return
		}
	}

	json.NewEncoder(w).Encode(&SpaSession{})
}
