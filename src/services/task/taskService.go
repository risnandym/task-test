package task_service

import (
	"task-test/src/contract"
	"task-test/src/entities"
)

func (p TaskService) Create(request contract.TaskInput) (response *entities.Task, err error) {

	task := &entities.Task{
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}

	response, err = p.taskRepo.Create(task)
	if err != nil {
		return
	}

	return
}
