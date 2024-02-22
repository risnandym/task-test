package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Task []Task `json:"-"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	hashedPassword, errPassword := u.GeneratePassword(u.Password)
	if errPassword != nil {
		return &User{}, errPassword
	}

	u.Password = string(hashedPassword)

	var err error = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) GeneratePassword(pass string) (string, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", errPassword
	}
	return string(hashedPassword), nil
}
