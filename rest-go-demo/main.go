package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


var db *gorm.DB
var err error

type Booking struct{
	Id      int    `json:”id”`
	User    string `json:”user”`
	Members int  `json:”members”`
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}


func handleRequests(){
	log.Println("Starting development server at http://127.0.0.1:10000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", returnAllBookings)
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking Booking
	json.Unmarshal(reqBody, &booking)
	db.Create(&booking)
	fmt.Println("Endpoint Hit: Creating New Booking")
	json.NewEncoder(w).Encode(booking)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request){
	bookings := []Booking{}
	db.Find(&bookings)
	fmt.Println("Endpoint Hit: returnAllBookings")
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
	db.Find(&bookings)
	for _, booking := range bookings{
		s, err := strconv.Atoi(key)
		if err == nil{
			if booking.Id == s {
				fmt.Println(booking)
				fmt.Println("Endpoint Hit: Booking No:", key)
				json.NewEncoder(w).Encode(booking)
			}
		}
	}
}



func main() {
	// Please define your username and password for MySQL.
	db, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/football?charset=utf8&parseTime=True")
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err!=nil{
		log.Println("Connection Failed to Open", err)
	}else{
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Booking{})
	handleRequests()


}