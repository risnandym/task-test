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
	response.ID = result.ID
	response.Name = result.Name
	response.Email = result.Email
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
			ID:        item.ID,
			Email:     item.Email,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}).ToSlice()

	return
}

func (u UserService) Update(id int, request contract.UpdateInput) (response *contract.RegisterOutput, err error) {

	result, err := u.userRepo.Get(id)
	if err != nil {
		return
	}

	user := &entities.User{
		ID:    result.ID,
		Name:  request.Name,
		Email: request.Email,
	}

	user, err = u.userRepo.Update(user)
	if err != nil {
		return
	}

	response = &contract.RegisterOutput{}
	response.ID = user.ID
	response.Email = user.Email
	response.Name = user.Name
	response.CreatedAt = user.CreatedAt
	response.UpdatedAt = user.UpdatedAt

	return
}

func (u UserService) Delete(id int) (err error) {
	return u.userRepo.Delete(id)
}
