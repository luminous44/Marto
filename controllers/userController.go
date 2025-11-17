package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/luminous44/Marto/config"
	"github.com/luminous44/Marto/models"
)

var jwtKey = []byte("vbnm123")

type Claims struct{
	Email string `json:email`
	jwt.StandardClaims
}

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

	var input models.User
	var user models.User
	json.NewDecoder(r.Body).Decode(&input)
	config.GetDB().Where("email = ?",input.Email).First(&user)

	if err := user.CheckPassword(input.Password); err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: input.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	tokenString, _ := token.SignedString(jwtKey)

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
	
     w.Write([]byte("Login successful"))
 
}