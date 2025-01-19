package models

type User struct{
	ID uint `gorm:"primaryKey"`
	Email string
	Username string
	Password string
}