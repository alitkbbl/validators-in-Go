package main

type User struct {
	FirstName   string `form:"first_name" binding:"required,min=2,max=16"`
	LastName    string `form:"last_name" binding:"required,min=2,max=16"`
	Username    string `form:"username" binding:"required,alphanum,min=8,max=32"`
	Email       string `form:"email" binding:"required,email,min=8,max=32"`
	PhoneNumber string `form:"phone_number" binding:"required,len=11,numeric,phonestart"`
	BirthDate   string `form:"birth_date" binding:"required,birthdate"`
	NationalID  string `form:"national_id" binding:"required,len=10,numeric,nationalid"`
}
