package handlers

import (
	"task-test/src/contract"
	"task-test/src/entities"
)

type UserService interface {
	Create(request contract.RegisterInput) (response *contract.RegisterOutput, err error)
	Login(request contract.LoginInput) (token string, err error)
	Get(id int) (response *contract.RegisterOutput, err error)
	GetList(request *contract.PageRequest) (response []contract.RegisterOutput, err error)
}

type TaskService interface {
	Create(request contract.TaskInput) (response *entities.Task, err error)
}
