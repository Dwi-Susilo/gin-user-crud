package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Con() (*gorm.DB, error)  {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

func Migrate(db *gorm.DB) error  {
	type User struct {
		ID       uint   `gorm:"primaryKey" json:"id"`
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"-"`
		CreatedAt time.Time `json:"created_at"`
    	UpdatedAt time.Time `json:"updated_at"`
	}
	return db.AutoMigrate(&User{})
}