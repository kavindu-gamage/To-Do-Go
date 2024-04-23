package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model

	TaskName    string `json:"name"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

//This gorm popular for object relational mapping
//used to simplify database operations by mapping go structs to datavse tables and providing methods to interact with database
