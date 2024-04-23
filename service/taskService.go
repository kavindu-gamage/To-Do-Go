package service

import (
	"go-todo/model"
	"go-todo/repository"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.TaskRepository.FindAllTasks()
}

func (s *TaskService) GetTaskById(id int) (*model.Task, error) {
	return s.TaskRepository.FindTaskByID(id)
}

func (s *TaskService) CreateTask(task *model.Task) error {
	return s.TaskRepository.SaveTask(task)
}

func (s *TaskService) UpdateTask(id int, task *model.Task) error {
	return s.TaskRepository.UpdateTask(id, task)
}

func (s *TaskService) DeleteTaskById(id int) error {
	return s.TaskRepository.DeleteTaskById(id)
}
