package models

import "time"

//Model campos base para todos os modelos
type Model struct {
	ID        string    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// User representa um usuário no sistema
type User struct {
	Model
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email" json:"email"`
}
