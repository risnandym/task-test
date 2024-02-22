package contract

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func ValidateAndBuildPageRequest(c *gin.Context) (request *PageRequest, err error) {

	pagestr := c.DefaultQuery("page", "1")
	limitstr := c.DefaultQuery("limit", "10")

	pageint, err := strconv.Atoi(pagestr)
	if err != nil {
		log.Printf("user login error: %v", err)
		return nil, err
	}

	limitint, err := strconv.Atoi(limitstr)
	if err != nil {
		log.Printf("user login error: %v", err)
		return nil, err
	}

	request = &PageRequest{
		Page:  pageint,
		Limit: limitint,
	}

	return
}
