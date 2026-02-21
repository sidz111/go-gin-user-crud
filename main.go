package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/user-management-crud/controller"
	"github.com/sidz111/user-management-crud/db"
	"github.com/sidz111/user-management-crud/repository"
	"github.com/sidz111/user-management-crud/service"
)

func main() {
	db := db.ConnectDb()

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Server is Running"})
	})

	users := router.Group("/users")
	{
		users.POST("/", userController.CreateUser)
		users.GET("/:id", userController.GetById)
		users.GET("/", userController.GetallUsers)
		users.DELETE("/:id", userController.DeleteById)
		users.PUT("/", userController.UpdateUser)
	}
	log.Println("Server is Started on Port 8080")
	router.Run(":8080")
}
