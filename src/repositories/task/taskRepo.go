package task_repo

import (
	"log"
	"task-test/src/contract"
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

func (t TaskRepository) Get(id int) (response entities.Task, err error) {

	result := t.db.First(&response, "id = ?", id)
	if result.Error != nil {
		log.Printf("get user error: %v", err)
		err = result.Error
		return
	}

	return
}

func (t TaskRepository) GetList(request *contract.PageRequest) (response []entities.Task, err error) {

	offset := (request.Page - 1) * request.Limit
	result := t.db.Offset(offset).Limit(request.Limit).Find(&response)
	if result.Error != nil {
		log.Printf("get user list error: %v", err)
		err = result.Error
		return
	}

	return
}

func (t TaskRepository) Update(request *entities.Task) (response *entities.Task, err error) {

	request.UpdatedAt = time.Now()

	result := t.db.Save(&request)
	if result.Error != nil {
		log.Printf("update user error: %v", err)
		err = result.Error
		return
	}

	response = request

	return
}

func (t TaskRepository) Delete(id int) (err error) {

	result := t.db.Delete(&entities.Task{}, id)
	if result.Error != nil {
		log.Printf("Delete user error: %v", err)
		err = result.Error
		return
	}

	return
}
