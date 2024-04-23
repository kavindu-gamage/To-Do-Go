package main

import (
	"go-todo/controller"
	"go-todo/db"
	"go-todo/repository"
	"go-todo/service"
	"go-todo/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitDB
}

func main() {
	r := gin.Default()

	userRepo := &repository.UserRepository{DB: db.GetDB()}
	userService := &service.UserService{UserRepository: userRepo}
	authController := &controller.AuthController{UserService: userService}

	taskRepo := &repository.TaskRepository{DB: db.GetDB()}
	taskService := &service.TaskService{TaskRepository: taskRepo}
	taskController := &controller.TaskController{TaskService: taskService}

	r.POST("/auth/register", authController.RegisterUser)
	r.POST("/auth/login", authController.UserLogin)

	authMiddleWare := utils.AuthMiddleWare()
	protected := r.Group("/")
	protected.Use(authMiddleWare)

	protected.GET("/tasks", taskController.GetTasksController)
	protected.GET("/task/:id", taskController.GetTaskController)
	protected.PUT("/task/:id", taskController.UpdateTaskController)
	protected.POST("/task", taskController.CreateTasksController)
	protected.DELETE("/task/:id", taskController.DeleteTaskController)

}
