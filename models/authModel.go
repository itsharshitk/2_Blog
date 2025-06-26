package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FName    string `json:"fname" validate:"required,max=50,min=2"`
	LName    string `json:"lname" validate:"required,max=50,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,validPass"`
}
