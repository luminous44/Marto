package controllers

import (
	"net/http"
	"github.com/luminous44/Marto/config"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    db := config.GetDB()
	if err := db.DB().Ping(); err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server and Database are healthy"))

}