package contract

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskOutput struct {
	TaskInput
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ValidateAndBuildTaskInput(c *gin.Context) (request TaskInput, err error) {
	userctx, _ := c.Get("user")
	user := userctx.(*RegisterOutput)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	log.Println(user.ID)

	request.UserID = user.ID
	return
}

func ValidateAndBuildUpdateTask(c *gin.Context) (id int, request TaskInput, err error) {

	id, err = GetQueryPathID(c)
	if err != nil {
		return
	}

	request, err = ValidateAndBuildTaskInput(c)
	if err != nil {
		return
	}

	return
}
