package task_service

import (
	"task-test/src/contract"
	"task-test/src/entities"
)

func (p TaskService) Create(request contract.TaskInput) (response *contract.TaskOutput, err error) {

	task := &entities.Task{
		UserID: int(request.UserID),
	}

	_, err = p.taskRepo.Create(task)
	if err != nil {
		return
	}

	response = &contract.TaskOutput{
		TaskInput: request,
	}
	return
}
