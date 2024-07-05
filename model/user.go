package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" validate:"required,username_unique"`
	Email    string `json:"email" validate:"required,email,email_unique"`
	Password string `json:"password,omitempty" validate:"required"`
}
