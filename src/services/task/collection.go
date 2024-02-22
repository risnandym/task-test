package task_service

import (
	"task-test/src/entities"

	"gorm.io/gorm"
)

type TaskService struct {
	db       *gorm.DB
	taskRepo TaskRepository
}

func NewTaskService(db *gorm.DB, taskRepo TaskRepository) *TaskService {
	return &TaskService{
		db:       db,
		taskRepo: taskRepo,
	}
}

type TaskRepository interface {
	Create(request *entities.Task) (response *entities.Task, err error)
}
