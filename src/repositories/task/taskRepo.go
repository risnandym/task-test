package task_repo

import (
	"task-test/src/entities"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) (*TaskRepository, error) {
	return &TaskRepository{
		db: db,
	}, nil
}

func (t TaskRepository) Create(request *entities.Task) (response *entities.Task, err error) {

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = t.db.Create(&request).Error; err != nil {
		return
	}

	response = request

	return
}
