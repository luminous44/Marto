package routes

import (
	"github.com/gorilla/mux"
	"github.com/luminous44/Marto/controllers"
)

var RegisterRoutes = func (router *mux.Router)  {

	router.HandleFunc("/health", controllers.HealthCheck ).Methods("GET")
	
}