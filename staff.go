package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Staffs struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func InitialMigration() {
	dns := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot establish connection with the database")
	}
	DB.AutoMigrate(&Staffs{})
}

func CreateStaffs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staffs Staffs
	json.NewDecoder(r.Body).Decode(&staffs) // decodes the request data
	DB.Create(&staffs)
	json.NewEncoder(w).Encode(staffs)
}

func GetStaffs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staffs []Staffs
	DB.Find(&staffs)
	json.NewEncoder(w).Encode(staffs)
}

func GetStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var staffs Staffs
	DB.First(&staffs, params["id"])
	json.NewEncoder(w).Encode(staffs)
}

func UpdateStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var staff Staffs
	DB.First(&staff, params["id"])
	json.NewDecoder(r.Body).Decode(&staff)
	DB.Save(&staff)
	json.NewEncoder(w).Encode(staff)
}

func DeleteStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var staff Staffs
	DB.Delete(&staff, params["id"])
	json.NewEncoder(w).Encode("Staff deleted successfully")
}
