package models

import (
	"github.com/jinzhu/gorm"
	"github.com/luminous44/Marto/config"
)

type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
}
var (
	db *gorm.DB
)
func init(){
	 config.Connection()
	 db = config.GetDB()	
	 db.AutoMigrate(Product{}) 
}

func CreateNewProduct(p *Product) *Product{
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAll() []Product{
	var allItem []Product
	db.Find(&allItem)
	return allItem
}

func GetByID(id int64) (*Product, *gorm.DB){
	var getItem Product
	db := db.Where("ID=?",id).Find(&getItem)
	return &getItem,db
}
func DeleteByID(id int64) Product {
    var product Product
    db.First(&product, id)
    if product.ID == 0 {
        return product
    }
    db.Delete(&product)
    return product
}
