package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dwi-Susilo/gin-user-crud/database"
	"github.com/Dwi-Susilo/gin-user-crud/routes"
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
	// log.Println("Database connected:", db != nil)

	r := gin.Default()
	routes.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}