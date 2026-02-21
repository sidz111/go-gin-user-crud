package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/user-management-crud/model"
	"github.com/sidz111/user-management-crud/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

// Insert
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	createdUser, err := c.service.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// Get by Id
func (c *UserController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "Invalid User ID"})
	}
	user, err := c.service.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	ctx.JSON(http.StatusOK, user)
}

// Delete By Id
func (c *UserController) DeleteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "Invalid id"})
	}
	if err := c.service.DeletebyId(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	ctx.JSON(http.StatusOK, "User deleted Successfull with id "+ctx.Param("id"))
}

// Get all Users
func (c *UserController) GetallUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if users == nil {
		users = []model.User{}
	}
	ctx.JSON(http.StatusOK, users)
}

// Update User
func (c *UserController) UpdateUser(ctx *gin.Context) {
	// id := strconv.Atoi(ctx.Param("id"))
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	if err := c.service.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, "User Updated with Id "+strconv.Itoa(user.ID))
}
