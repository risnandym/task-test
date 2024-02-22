package user_service

import (
	"task-test/src/contract"
	"task-test/src/entities"

	"github.com/mariomac/gostream/stream"
)

func (u UserService) Create(request contract.RegisterInput) (response *contract.RegisterOutput, err error) {

	user := entities.User{}
	user.Email = request.Email
	user.Name = request.Name
	user.Password = request.Password

	result, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	response = &contract.RegisterOutput{}
	response.Email = result.Email
	response.Name = result.Name
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt

	return
}

func (u UserService) Login(request contract.LoginInput) (token string, err error) {

	user := entities.User{}
	user.Email = request.Email
	user.Password = request.Password
	token, err = u.userRepo.Login(user.Email, user.Password)

	return
}

func (u UserService) Get(id int) (response *contract.RegisterOutput, err error) {
	result, err := u.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	response = &contract.RegisterOutput{}
	response.ID = result.ID
	response.Email = result.Email
	response.Name = result.Name
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt

	return
}

func (u UserService) GetList(request *contract.PageRequest) (response []contract.RegisterOutput, err error) {

	users, err := u.userRepo.GetList(request)
	if err != nil {
		return nil, err
	}

	response = stream.Map(stream.OfSlice(users), func(item entities.User) contract.RegisterOutput {
		return contract.RegisterOutput{
			Email:     item.Email,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}).ToSlice()

	return
}
