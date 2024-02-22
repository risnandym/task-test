package task_service

import (
	"task-test/src/contract"
	"task-test/src/entities"

	"github.com/mariomac/gostream/stream"
)

func (t TaskService) Create(request contract.TaskInput) (response *entities.Task, err error) {

	task := &entities.Task{
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}

	response, err = t.taskRepo.Create(task)
	if err != nil {
		return
	}

	return
}

func (t TaskService) GetList(request *contract.PageRequest) (response []entities.Task, err error) {

	tasks, err := t.taskRepo.GetList(request)
	if err != nil {
		return nil, err
	}

	response = stream.Map(stream.OfSlice(tasks), func(item entities.Task) entities.Task {
		return entities.Task{
			ID:          item.ID,
			UserID:      item.UserID,
			Title:       item.Title,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}).ToSlice()

	return
}

func (t TaskService) Get(id int) (response *entities.Task, err error) {

	result, err := t.taskRepo.Get(id)
	if err != nil {
		return nil, err
	}

	response = &entities.Task{}
	response.ID = result.ID
	response.UserID = result.UserID
	response.Description = result.Description
	response.Status = result.Status
	response.Title = result.Title
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt

	return
}

func (t TaskService) Update(id int, request contract.TaskInput) (response *entities.Task, err error) {

	result, err := t.taskRepo.Get(id)
	if err != nil {
		return nil, err
	}

	task := &entities.Task{
		ID:          result.ID,
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}

	task, err = t.taskRepo.Update(task)
	if err != nil {
		return nil, err
	}

	response = &entities.Task{}
	response.ID = task.ID
	response.UserID = task.UserID
	response.Description = task.Description
	response.Status = task.Status
	response.Title = task.Title
	response.CreatedAt = task.CreatedAt
	response.UpdatedAt = task.UpdatedAt

	return
}

func (t TaskService) Delete(id int) (err error) {
	return t.taskRepo.Delete(id)
}
