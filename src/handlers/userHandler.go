package handlers

import (
	"net/http"
	"task-test/src/contract"

	"github.com/gin-gonic/gin"
)

type ChangePasswordInput struct {
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"phone"`
	Password string `json:"password" binding:"required"`
}

// LoginUser godoc
//
//	@Summary		Login.
//	@Description	Logging in to get jwt token to access admin or user api by roles.
//	@Tags			User
//	@Param			Body	body	contract.LoginInput	true	"the body to login a user"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/task-test/login [post]
func Login(svc UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildUserLogin(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Login(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password is incorrect."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "login success", "token": response})
	}
}

// Register godoc
//
//	@Summary		Register a User.
//	@Description	registering a user from public access.
//	@Tags			User
//	@Param			Body	body	contract.RegisterInput	true	"the body to register a user"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/task-test/users [post]
func Register(svc UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildUserRegister(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Create(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": response})
	}
}

// GetList godoc
//
//	@Summary		Get List User.
//	@Description	get list user from public access.
//	@Tags			User
//	@Security		task-token
//	@Param			page	query	string	false	"page"
//	@Param			limit	query	string	false	"limit"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/task-test/users [get]
func GetList(svc UserService) gin.HandlerFunc {
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
