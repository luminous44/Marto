package models
import(
    "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	gorm.Model
	Name string `json: "name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func(u *User) HashPassword(pass string) (string, error){ 
  bytes, err := bcrypt.GenerateFromPassword([]byte(pass),14)
  return string(bytes),err
}
