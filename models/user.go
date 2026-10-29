package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name" binding:"required"`
    Email     string    `json:"email" binding:"required,email"`
    Password  string    `json:"password,omitempty" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}