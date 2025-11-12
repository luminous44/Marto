package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luminous44/Marto/config"
	"github.com/luminous44/Marto/routes"
)

func main() {
    config.Connection()
	r := mux.NewRouter()
    routes.RegisterRoutes(r)
	fmt.Println("server running on server :8080")
	if err := http.ListenAndServe(":8080",r); err != nil {
		fmt.Println("Failed to start server",err)  
		log.Fatal(err)
		
	}
}
