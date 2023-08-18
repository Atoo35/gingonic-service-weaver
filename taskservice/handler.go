package taskservice

import (
	"context"
	"net/http"
	"strconv"

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

	body.ID = strconv.Itoa(len(mock.Tasks) + 1)
	alltasks := append(mock.Tasks, body)
	gctx.JSON(http.StatusCreated, gin.H{
		"tasks": alltasks,
	})
}

func (s *Server) UpdateTask(gctx *gin.Context) {
	id := gctx.Param("id")
	task := getTaskByID(id)

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

	var result []models.Task
	for _, t := range mock.Tasks {
		if t.ID == id {
			result = append(result, body)
		} else {
			result = append(result, t)
		}
	}

	gctx.JSON(http.StatusCreated, gin.H{
		"tasks": result,
	})
}

func (s *Server) DeleteTask(gctx *gin.Context) {
	id := gctx.Param("id")
	task := getTaskByID(id)

	if task.ID == "" {
		gctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
		return
	}

	var result []models.Task
	for _, t := range mock.Tasks {
		if t.ID != id {
			result = append(result, t)
		}
	}

	gctx.JSON(http.StatusCreated, gin.H{
		"tasks": result,
	})
}
