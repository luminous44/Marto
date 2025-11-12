package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luminous44/Marto/routes"
)

func main() {
	r := mux.NewRouter()
    routes.RegisterRoutes(r)
	fmt.Println("server running on server :9000")
	if err := http.ListenAndServe(":9000",r); err != nil {
		fmt.Println("Failed to start server",err)  
		log.Fatal(err)
		
	}
}
