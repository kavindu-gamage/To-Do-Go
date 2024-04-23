package controller

import (
	"net/http"
	"strconv"

	"example.com/hello/Documents/SE-Projects/go-todo/model"
	"example.com/hello/Documents/SE-Projects/go-todo/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TaskController struct {
	TaskService *service.TaskService
}

func (h *TaskController) GetTasksController(c *gin.Context) {
	tasks, err := h.TaskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retriving Tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskController) GetTaskController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	task, err := h.TaskService.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retriving Task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskController) CreateTaskController(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.TaskService.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskController) UpdateTaskController(c *gin.Context) {
	var updatedTask model.Task
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err = h.TaskService.UpdateTask(id, &updatedTask)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating Task"})
		}
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskController) DeleteTaskController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.TaskService.DeleteTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Deleting Task"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
