package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"gte=0,lte=130"`
}
