package routes

import (
	"github.com/Atoo35/gingonic-service-weaver/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	h := &handlers.TaskHandler{}
	router := gin.Default()

	tasksRoutes := router.Group("/api/tasks")
	{
		tasksRoutes.GET("/", h.GetTasks)
		tasksRoutes.POST("/", h.CreateTask)
		tasksRoutes.GET("/:id", h.GetTask)
		// tasksRoutes.PUT("/:id", h.UpdateTask)
		// tasksRoutes.DELETE("/:id", h.DeleteTask)
	}

	return router
}
