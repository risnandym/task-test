package middlewares

import (
	"log"
	"net/http"
	"task-test/core/config"
	"task-test/core/utils"
	"task-test/src/handlers"

	"github.com/gin-gonic/gin"
)

func PublicMiddleware(svc handlers.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := utils.TokenValid(c, config.Config())
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		id, _ := utils.ExtractTokenID(c, config.Config())

		user, err := svc.Get(id)
		if err != nil {
			log.Fatal(err)
		}

		c.Set("user", user)

		c.Next()
	}
}
