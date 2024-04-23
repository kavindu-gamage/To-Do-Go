package repository

import (
	"go-todo/model"

	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (r *TaskRepository) FindTaskByID(id int) (*model.Task, error) {
	task := model.Task{}
	err := r.DB.First(&task, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) FindAllTask() ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.DB.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (r *TaskRepository) SaveTask(task *model.Task) error {
	err := r.DB.Create(task).Error
	return err
}

func (r *TaskRepository) UpdateTask(id int, task *model.Task) error {
	err := r.DB.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
	return err
}

func (r *TaskRepository) DeleteTaskById(id int) error {
	err := r.DB.Delete(&model.Task{}, id).Error
	return err
}
