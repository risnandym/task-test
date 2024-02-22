package contract

import (
	"time"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterOutput struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ValidateAndBuildUserRegister(c *gin.Context) (request RegisterInput, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}
	return
}

func ValidateAndBuildUserLogin(c *gin.Context) (request LoginInput, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}
	return
}
