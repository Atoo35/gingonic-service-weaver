package taskservice

import (
	"context"

	"github.com/Atoo35/gingonic-service-weaver/notificationservice"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-gonic/gin"
)

type Server struct {
	weaver.Implements[weaver.Main]
	taskapi weaver.Listener
	handler *gin.Engine

	notificationService weaver.Ref[notificationservice.Service]
}

func (s *Server) Init(ctx context.Context) error {
	router := gin.Default()

	tasksRoutes := router.Group("/api/tasks")
	{
		tasksRoutes.GET("/", s.GetTasks)
		tasksRoutes.POST("/", s.CreateTask)
		tasksRoutes.GET("/:id", s.GetTask)
		tasksRoutes.PUT("/:id", s.UpdateTask)
		tasksRoutes.DELETE("/:id", s.DeleteTask)
	}
	s.handler = router
	return nil
}

func Serve(ctx context.Context, server *Server) error {
	server.Logger(ctx).Info("Task API listening on ", "addr:", server.taskapi)
	router := server.handler
	return router.RunListener(server.taskapi)
}
