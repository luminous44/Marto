package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/luminous44/Marto/config"
	"github.com/luminous44/Marto/models"
)

func Register(w http.ResponseWriter, r *http.Request){
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	hashPass, err := newUser.HashPassword(newUser.Password)
	if(err != nil){
		return
	}
	newUser.Password = hashPass
	config.GetDB().Create(&newUser)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUser)

}
func Login(w http.ResponseWriter, r *http.Request){

}