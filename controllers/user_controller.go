package controllers

import (
	"errors"
	"net/http"
	"technical-test/models"
	"technical-test/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func jsonResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func jsonError(c *gin.Context, status int, message string, err string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"error":   err,
	})
}

func (ctrl *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}
	if err := ctrl.UserService.Register(&user); err != nil {
		if err.Error() == "username or email already exists" {
			jsonError(c, http.StatusConflict, "Username or email already exists", err.Error())
			return
		}
		jsonError(c, http.StatusInternalServerError, "Registration failed", err.Error())
		return
	}

	jsonResponse(c, http.StatusOK, "Registration successful", nil)
}


func (ctrl *UserController) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}
	token, err := ctrl.UserService.Login(credentials.Username, credentials.Password)
	if err != nil {
		jsonError(c, http.StatusUnauthorized, "Authentication failed", err.Error())
		return
	}
	jsonResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

func (ctrl *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := ctrl.UserService.GetUserByUsername(username)
	if err != nil {
		jsonError(c, http.StatusNotFound, "User not found", "No user with the specified username")
		return
	}
	jsonResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.UserService.GetAllUsers()
	if err != nil {
		jsonError(c, http.StatusInternalServerError, "Failed to retrieve users", err.Error())
		return
	}
	jsonResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	username := c.Param("username")
	var updateData struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid request", "Failed to bind JSON")
		return
	}
	err := ctrl.UserService.UpdateUser(username, updateData.Email, updateData.Username)
	if err != nil {
		jsonError(c, http.StatusInternalServerError, "Update failed", err.Error())
		return
	}
	jsonResponse(c, http.StatusOK, "User updated successfully", nil)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	if err := ctrl.UserService.DeleteUserByUsername(username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsonError(c, http.StatusNotFound, "User not found", "No user with the specified username")
			return
		}
		jsonError(c, http.StatusInternalServerError, "Deletion failed", err.Error())
		return
	}
	jsonResponse(c, http.StatusOK, "User deleted successfully", nil)
}
