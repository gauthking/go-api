package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/staff", GetStaffs).Methods("GET")
	r.HandleFunc("/staff/{id}", GetStaff).Methods("GET")
	r.HandleFunc("/staff", CreateStaffs).Methods("POST")
	r.HandleFunc("/staff/{id}", UpdateStaff).Methods("PUT")
	r.HandleFunc("/staff/{id}", DeleteStaff).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	InitialMigration()
	initializeRouter()
}
