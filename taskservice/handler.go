package taskservice

import (
	"context"
	"net/http"

	"github.com/Atoo35/gingonic-service-weaver/mock"
	"github.com/Atoo35/gingonic-service-weaver/models"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func (s *Server) GetTasks(gctx *gin.Context) {
	tasks := mock.Tasks
	if err := s.notificationService.Get().Send(ctx); err != nil {
		s.Logger(ctx).Error("Failed to send notif")
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while sending notification",
		})
	}
	gctx.JSON(http.StatusAccepted, gin.H{
		"tasks": tasks,
	})
}

func (s *Server) GetTask(gctx *gin.Context) {
	id := gctx.Param("id")
	task := new(models.Task)
	for _, value := range mock.Tasks {
		if value.ID == id {
			task = &value
			break
		}
	}

	if task.ID != "" {
		gctx.JSON(http.StatusOK, gin.H{
			"task": task,
		})
	} else {
		gctx.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
	}
}

// func (t *TaskHandler) CreateTask(gctx *gin.Context)
