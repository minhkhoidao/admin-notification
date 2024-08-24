package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Role     string `json:"role" gorm:"not null;default:admin"`
}
