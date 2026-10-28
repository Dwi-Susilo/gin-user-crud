package main

import (
	"log"

	"github.com/Dwi-Susilo/gin-user-crud/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main()  {
	db, err := database.Con()
	
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	log.Println("Database connected:", db != nil)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}