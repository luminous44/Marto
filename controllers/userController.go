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
    Role  string `json:"role"`
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
func Login(w http.ResponseWriter, r *http.Request) {
    // Parse input JSON
    var input models.User
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Find user by email
    var user models.User
    result := config.GetDB().Where("email = ?", input.Email).First(&user)
    if result.Error != nil {
        http.Error(w, "Email not found", http.StatusUnauthorized)
        return
    }

    // Check password
    if err := user.CheckPassword(input.Password); err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Create JWT token
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        Email: user.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Use HS256 because you are signing with a []byte secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Token generation failed", http.StatusInternalServerError)
        return
    }

    // Set token as cookie (optional)
    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    tokenString,
        Expires:  expirationTime,
    })

    // Return token in JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}

 
