package handlers

import (
	"net/http"
	"task-test/src/contract"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
//
//	@Summary		Create Task.
//	@Description	Save User Task.
//	@Tags			Task
//	@Param			Body	body	contract.TaskInput	true	"the body to create a new Task"
//	@Security		task-token
//	@Produce		json
//	@Success		200	{object}	contract.TaskOutput
//	@Router			/task-test/tasks [post]
func CreateTask(svc TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildTaskInput(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Create(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}
