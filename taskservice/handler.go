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
	tasks, err := s.taskRepository.Get().GetTasks(ctx)
	if err != nil {
		s.Logger(ctx).Error("Failed to get tasks", err)
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong getting tasks",
			"error":   err,
		})
	}
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
	task, err := s.taskRepository.Get().GetTaskByID(ctx, id)
	if err != nil {
		s.Logger(ctx).Error("Failed to get task by id: %s", id, err)
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong getting task",
			"error":   err,
		})
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

func getTaskByID(id string) *models.Task {
	task := new(models.Task)
	for _, value := range mock.Tasks {
		if value.ID == id {
			task = &value
			break
		}
	}
	return task
}

func (s *Server) CreateTask(gctx *gin.Context) {
	body := models.Task{}

	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad body",
		})
		return
	}

	err := s.taskRepository.Get().CreateTask(ctx, body)
	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while creating task",
			"error":   err,
		})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{
		"task": body,
	})
}

func (s *Server) UpdateTask(gctx *gin.Context) {
	id := gctx.Param("id")
	task, err := s.taskRepository.Get().GetTaskByID(ctx, id)

	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"error":   err,
		})
		return
	}
	if task.ID == "" {
		gctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
		return
	}

	body := models.Task{}

	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad body",
		})
		return
	}

	updatedTask, err := s.taskRepository.Get().UpdateTask(ctx, id, body)
	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"error":   err,
		})
		return
	}

	gctx.JSON(http.StatusCreated, gin.H{
		"tasks": updatedTask,
	})
}

func (s *Server) DeleteTask(gctx *gin.Context) {
	id := gctx.Param("id")

	if err := s.taskRepository.Get().DeleteTask(ctx, id); err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while deleting the task",
			"error":   err,
		})
	}

	gctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted.",
	})
}
