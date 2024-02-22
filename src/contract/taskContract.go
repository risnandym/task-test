package contract

import (
	"task-test/src/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	UserID       int       `json:"user_id,omitempty"`
	NIK          string    `json:"nik" validate:"required"`
	FullName     string    `json:"full_name" validate:"required"`
	LegalName    string    `json:"legal_name" validate:"required"`
	PlaceOfBirth string    `json:"place_of_birth" validate:"required"`
	DateOfBirth  time.Time `json:"date_of_birth" validate:"required"`
	Salary       float64   `json:"salary" validate:"required"`
	KtpImage     string    `json:"ktp_image" validate:"required"`
	SelfieImage  string    `json:"selfie_image" validate:"required"`
}

type TaskOutput struct {
	TaskInput
}

func ValidateAndBuildTaskInput(c *gin.Context) (request TaskInput, err error) {
	userctx, _ := c.Get("user")
	user := userctx.(entities.User)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	request.UserID = user.ID
	return
}
