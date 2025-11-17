package models
import(
    "github.com/jinzhu/gorm"
)

type User struct{
	gorm.DB

	Name string `json: "name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}