package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"spa_service/booking"
)

type SpaSession struct {
	Id          string `json:"Id"`
	Session     string `json:"Session"`
	Date        string `json:"Date"`
	Time        string `json:"Time"`
	Description string `json:"Description"`
}

var SpaSessions []SpaSession

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

func returnAllBookings(w http.ResponseWriter, r *http.Request){
	fmt.Println("returnAllBookings")
	json.NewEncoder(w).Encode(booking.Bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["uid"]
	for _, nbooking := range booking.Bookings {
		if nbooking.Uid == key {
			json.NewEncoder(w).Encode(nbooking)
		}
	}
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var nbooking booking.Booking
	json.Unmarshal(reqBody, &nbooking)
	booking.Bookings = append(booking.Bookings, nbooking)
	json.NewEncoder(w).Encode(nbooking)
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uid"]
	for index, nbooking := range booking.Bookings {
		if nbooking.Uid == key {
			booking.Bookings = append(booking.Bookings[:index], booking.Bookings[index+1:]...)
		}
	}
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uid"]
	var tempBooking booking.Booking
	for index, nbooking := range booking.Bookings {
		if (nbooking.Uid) == key {
			_ = json.NewDecoder(r.Body).Decode(&tempBooking)
			tempBooking.Uid = strconv.Itoa(index)
			booking.Bookings[index] = tempBooking
			json.NewEncoder(w).Encode(booking.Bookings[index])
			return
		}
	}

	json.NewEncoder(w).Encode(&SpaSession{})
}

func returnSessionBooking(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	for _, nbooking := range booking.Bookings {
		if nbooking.SessionID == key {
			json.NewEncoder(w).Encode(nbooking)
		}
	}
}
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
	r.HandleFunc("/bookings", returnAllBookings)
	r.HandleFunc("/booking", createBooking).Methods("POST")
	r.HandleFunc("/booking/{uid}", deleteBooking).Methods("DELETE")
	r.HandleFunc("/booking/{uid}", updateBooking).Methods("PUT")
	r.HandleFunc("/booking/{uid}", returnSingleBooking)
	r.HandleFunc("/bookingSession/{id}", returnSessionBooking)
	log.Fatal(http.ListenAndServe(":10000", r))

}

func main() {
	fmt.Println("Welcome to Spa Service")
	handleRequest()
}


