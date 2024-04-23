package controller

import (
	"go-todo/model"
	"go-todo/service"
	"go-todo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *service.UserService
}

func (c *AuthController) RegisterUser(cxt *gin.Context) {
	var user model.User
	err := cxt.ShouldBindJSON(&user)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data"})
		return
	}
	err = c.UserService.RegisterUser(&user)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "User Registered Succesfully"})
}

func (c *AuthController) UserLogin(cxt *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := cxt.ShouldBindJSON(&loginData)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data"})
		return
	}
	user, err := c.UserService.AuthenticateUser(loginData.Username, loginData.Password)
	if err != nil {
		cxt.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	cxt.JSON(http.StatusOK, gin.H{"token": token})
}
