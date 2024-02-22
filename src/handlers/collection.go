package handlers

import (
	"task-test/src/contract"
)

type UserService interface {
	Create(request contract.RegisterInput) (response *contract.RegisterOutput, err error)
	Login(request contract.LoginInput) (token string, err error)
	Get(id uint) (response *contract.RegisterOutput, err error)
	GetList(request *contract.PageRequest) (response []contract.RegisterOutput, err error)
}

type TaskService interface {
	Create(request contract.TaskInput) (response *contract.TaskOutput, err error)
}
