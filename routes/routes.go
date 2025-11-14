package routes

import (
	"github.com/gorilla/mux"
	"github.com/luminous44/Marto/controllers"
)

var RegisterRoutes = func (r *mux.Router)  {

	r.HandleFunc("/health", controllers.HealthCheck ).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
    r.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
    r.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
    // r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	
}