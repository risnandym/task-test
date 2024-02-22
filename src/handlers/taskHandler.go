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

// GetListTask godoc
//
//	@Summary		Get List Task.
//	@Description	get list task.
//	@Tags			Task
//	@Security		task-token
//	@Param			page	query	string	false	"page"
//	@Param			limit	query	string	false	"limit"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/task-test/tasks [get]
func GetListTask(svc TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildPageRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.GetList(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": response})
	}
}

// Get Task godoc
//
//	@Summary		Get a Task.
//	@Description	get a task.
//	@Tags			Task
//	@Security		task-token
//	@Param			id	path	int	true	"task id"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/task-test/tasks/{id} [get]
func GetTask(svc TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.GetQueryPathID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Get(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": response})
	}
}

// UpdateTask godoc
//
//	@Summary		Update Task.
//	@Description	Edit User Task.
//	@Tags			Task
//	@Param			id		path	int					true	"task id"
//	@Param			Body	body	contract.TaskInput	true	"the body to create a new Task"
//	@Security		task-token
//	@Produce		json
//	@Success		200	{object}	contract.TaskOutput
//	@Router			/task-test/tasks/{id} [put]
func UpdateTask(svc TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, request, err := contract.ValidateAndBuildUpdateTask(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Update(id, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}

// DeleteTask godoc
//
//	@Summary		Delete a Task.
//	@Description	remove a task.
//	@Tags			Task
//	@Security		task-token
//	@Param			id	path	int	true	"task id"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/task-test/tasks/{id} [delete]
func DeleteTask(svc TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.GetQueryPathID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = svc.Delete(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete success"})
	}
}
