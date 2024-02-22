package src

import (
	"net/http"
	"task-test/core/middlewares"
	"task-test/src/handlers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(app *Dependency) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	base := r.Group("/task-test")
	base.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	base.POST("/login", handlers.Login(app.Services.UserSVC))

	user := base.Group("/users")
	user.POST("/", handlers.Register(app.Services.UserSVC))

	user.Use(middlewares.PublicMiddleware(app.Services.UserSVC))
	user.GET("/", handlers.GetList(app.Services.UserSVC))
	// user.GET("/:id", handlers.Get(app.Services.UserSVC))
	// user.PUT("/:id", handlers.Update(app.Services.UserSVC))
	// user.DELETE("/:id", handlers.Delete(app.Services.UserSVC))

	task := base.Group("/tasks")
	task.POST("/", handlers.CreateTask(app.Services.TaskSVC))
	// task.GET("/", handlers.GetList(app.Services.TaskSVC))

	// task.Use(middlewares.PublicMiddleware(app.Services.TaskSVC))
	// task.GET("/:id", handlers.Get(app.Services.TaskSVC))
	// task.PUT("/:id", handlers.Update(app.Services.TaskSVC))
	// task.DELETE("/:id", handlers.Delete(app.Services.TaskSVC))

	return r
}
