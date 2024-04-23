package main

import (
	"example.com/hello/Documents/SE-Projects/go-todo/controller"
	"example.com/hello/Documents/SE-Projects/go-todo/db"
	"example.com/hello/Documents/SE-Projects/go-todo/repository"
	"example.com/hello/Documents/SE-Projects/go-todo/service"
	"example.com/hello/Documents/SE-Projects/go-todo/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitDB()
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
	protected.POST("/task", taskController.CreateTaskController)
	protected.DELETE("/task/:id", taskController.DeleteTaskController)

	r.Run(":9000")

}
