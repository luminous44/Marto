package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luminous44/Marto/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request){

	var product  models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	_ = *models.CreateNewProduct(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	
}
func GetAllProducts(w http.ResponseWriter, r *http.Request){
	all := models.GetAll()
	res, _ := json.Marshal(all)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetProductByID(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	productId := vars["id"]

	Id, err := strconv.ParseInt(productId,0,0)
	 if err != nil{
		fmt.Println("error while parsing")
	}
	productDetails, _ := models.GetByID(Id)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	productId := vars["id"]

	Id, err := strconv.ParseInt(productId,0,0)
	 if err != nil{
		fmt.Println("error while parsing")
	}
	productDetails := models.DeleteByID(Id)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}