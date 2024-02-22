package user_service

import (
	"task-test/src/contract"
	"task-test/src/entities"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type UserRepository interface {
	Create(request entities.User) (response *entities.User, err error)
	Login(username string, password string) (token string, err error)
	Get(id uint) (response entities.User, err error)
	GetList(request *contract.PageRequest) (response []entities.User, err error)
}
