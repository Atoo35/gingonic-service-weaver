package taskservice

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/gin-gonic/gin"
)

type Server struct {
	weaver.Implements[weaver.Main]
	taskapi weaver.Listener
	handler *gin.Engine
}

func (s *Server) Init(ctx context.Context) error {
	h := &TaskHandler{}
	router := gin.Default()

	tasksRoutes := router.Group("/api/tasks")
	{
		tasksRoutes.GET("/", h.GetTasks)
		// tasksRoutes.POST("/", h.CreateTask)
		tasksRoutes.GET("/:id", h.GetTask)
		// tasksRoutes.PUT("/:id", h.UpdateTask)
		// tasksRoutes.DELETE("/:id", h.DeleteTask)
	}
	s.handler = router
	return nil
}

func Serve(ctx context.Context, server *Server) error {
	server.Logger(ctx).Info("Task API listening on ", "addr:", server.taskapi)
	router := server.handler
	return router.RunListener(server.taskapi)
}
