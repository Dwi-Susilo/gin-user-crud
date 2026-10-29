package routes

import (
	"github.com/Dwi-Susilo/gin-user-crud/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db interface{})  {
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/api/v1") 
	{
		users := v1.Group("/users")
		{
			users.POST("", controllers.CreateUser)
			users.GET("", controllers.GetUsers)
			users.GET("/:id", controllers.GetUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}


}