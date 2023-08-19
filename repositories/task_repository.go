package repositories

import (
	"context"
	"errors"

	"github.com/Atoo35/gingonic-service-weaver/models"
	"github.com/ServiceWeaver/weaver"
	"github.com/surrealdb/surrealdb.go"
)

type config struct {
	User       string
	Pass       string
	ConnString string
}

type TaskRepository interface {
	GetTasks(context.Context) ([]models.Task, error)
	GetTaskByID(context.Context, string) (models.Task, error)
	CreateTask(context.Context, models.Task) error
	DeleteTask(context.Context, string) error
	UpdateTask(context.Context, string, models.Task) (models.Task, error)
}

type taskRepository struct {
	weaver.Implements[TaskRepository]
	weaver.WithConfig[config]
	db *surrealdb.DB
}

func (cfg *config) Validate() error {
	if len(cfg.User) == 0 {
		return errors.New("DB user is not provided")
	}
	if len(cfg.Pass) == 0 {
		return errors.New("DB pass is not provided")
	}
	if len(cfg.ConnString) == 0 {
		return errors.New("DB connstring is not provided")
	}
	return nil
}

func (taskRepo *taskRepository) Init(ctx context.Context) error {
	cfg := taskRepo.Config()
	var err error
	logger := taskRepo.Logger(ctx)
	if err := cfg.Validate(); err != nil {
		logger.Error("error:", err)
		return err
	}

	taskRepo.db, err = surrealdb.New(cfg.ConnString)
	if err != nil {
		logger.Error("Failed to connect", err)
		return err
	}
	_, err = taskRepo.db.Signin(map[string]interface{}{
		"user": cfg.User,
		"pass": cfg.Pass,
	})

	if err != nil {
		logger.Error("Error signing in: %s", err)
		return err
	}

	if _, err = taskRepo.db.Use("test", "tasks"); err != nil {
		logger.Error("Error using database: %s", err)
		return err
	}

	logger.Info("Connected to db with namespace test and collection tasks")
	return nil
}

func (taskRepo *taskRepository) getTaskByID(id string) (interface{}, error) {
	task, err := taskRepo.db.Select(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (taskRepo *taskRepository) GetTasks(ctx context.Context) ([]models.Task, error) {
	logger := taskRepo.Logger(ctx)
	tasks, err := taskRepo.db.Select("tasks")
	if err != nil {
		logger.Error("Error while selecting tasks: %s", err)
		return nil, err
	}

	tasksSlice := new([]models.Task)
	err = surrealdb.Unmarshal(tasks, &tasksSlice)
	if err != nil {
		logger.Error("Error while unmarshalling", err)
		return nil, err
	}
	return *tasksSlice, nil
}

func (taskRepo *taskRepository) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	logger := taskRepo.Logger(ctx)
	task, err := taskRepo.getTaskByID(id)

	if err != nil {
		logger.Error("Error while selecting tasks: %s", err)
		return models.Task{}, err
	}
	taskModel := new(models.Task)
	err = surrealdb.Unmarshal(task, &taskModel)
	if err != nil {
		logger.Error("Error while unmarshalling", err)
		return models.Task{}, err
	}

	return *taskModel, nil
}

func (taskRepo *taskRepository) CreateTask(ctx context.Context, task models.Task) error {
	logger := taskRepo.Logger(ctx)
	_, err := taskRepo.db.Create("tasks", task)
	if err != nil {
		logger.Error("Error while creating task", err)
		return err
	}
	return nil
}

func (taskRepo *taskRepository) DeleteTask(ctx context.Context, id string) error {
	logger := taskRepo.Logger(ctx)
	_, err := taskRepo.db.Delete(id)

	if err != nil {
		logger.Error("Error while deleting task", err)
		return err
	}

	return nil
}

func (taskRepo *taskRepository) UpdateTask(ctx context.Context, id string, task models.Task) (models.Task, error) {
	logger := taskRepo.Logger(ctx)
	_, err := taskRepo.db.Update(id, task)

	if err != nil {
		logger.Error("Error while updating task", err)
		return models.Task{}, err
	}

	return task, nil
}
