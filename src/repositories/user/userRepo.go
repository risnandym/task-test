package user_repo

import (
	"log"
	"task-test/core/config"
	"task-test/core/utils"
	"task-test/src/contract"
	"task-test/src/entities"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	config config.Configuration
}

func NewUserRepository(db *gorm.DB, config config.Configuration) (*UserRepository, error) {
	return &UserRepository{
		db:     db,
		config: config,
	}, nil
}

func (u UserRepository) Create(request entities.User) (response *entities.User, err error) {
	log.Println(u.db)
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	response, err = request.SaveUser(u.db)
	if err != nil {
		log.Printf("user create error: %v", err)
		return
	}

	return
}

func (u UserRepository) Login(email string, password string) (token string, err error) {

	user := entities.User{}

	err = u.db.Model(entities.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		log.Printf("user login error: %v", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Printf("user login error: %v", err)
		return "", err
	}

	token, err = utils.GenerateToken(u.config, user.ID)
	if err != nil {
		log.Printf("user login error: %v", err)
		return "", err
	}

	return token, nil
}

func (u UserRepository) Get(id int) (response entities.User, err error) {

	result := u.db.First(&response, "id = ?", id)
	if result.Error != nil {
		log.Printf("get user error: %v", err)
		err = result.Error
		return
	}

	return
}

func (u UserRepository) GetList(request *contract.PageRequest) (response []entities.User, err error) {

	offset := (request.Page - 1) * request.Limit
	result := u.db.Offset(offset).Limit(request.Limit).Find(&response)
	if result.Error != nil {
		log.Printf("get user list error: %v", err)
		err = result.Error
		return
	}

	return
}

func (u UserRepository) Update(request *entities.User) (response *entities.User, err error) {

	request.UpdatedAt = time.Now()

	result := u.db.Save(&request)
	if result.Error != nil {
		log.Printf("update user error: %v", err)
		err = result.Error
		return
	}

	response = request

	return
}

func (u UserRepository) Delete(id int) (err error) {

	result := u.db.Delete(&entities.User{}, id)
	if result.Error != nil {
		log.Printf("update user error: %v", err)
		err = result.Error
		return
	}

	return
}
