package handlers

import (
	"net/http"
	"strconv"

	"github.com/Atoo35/gingonic-service-weaver/mock"
	"github.com/Atoo35/gingonic-service-weaver/models"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
}

func (t *TaskHandler) GetTasks(gctx *gin.Context) {
	tasks := mock.Tasks
	gctx.JSON(http.StatusAccepted, gin.H{
		"tasks": tasks,
	})
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

func (t *TaskHandler) GetTask(gctx *gin.Context) {
	id := gctx.Param("id")
	task := getTaskByID(id)
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

func (t *TaskHandler) CreateTask(gctx *gin.Context) {
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

func (t *TaskHandler) UpdateTask(gctx *gin.Context) {
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

func (t *TaskHandler) DeleteTask(gctx *gin.Context) {
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
