package services

import (
	"context"
	"errors"
	"sync"

	"github.com/oxxi/jobsity/models"
	"github.com/oxxi/jobsity/repositories"
)

type ITaskService interface {
	GetAllTask(ctx context.Context) ([]models.Task, error)
	GetById(ctx context.Context, id int) (models.Task, error)
	Save(ctx context.Context, model models.Task) (models.Task, error)
	Update(ctx context.Context, id int, model models.Task) (models.Task, error)
	Delete(ctx context.Context, id int) error
}

type taskService struct {
	repo repositories.ITaskRepository
}

// Delete implements ITaskService.
func (t *taskService) Delete(ctx context.Context, id int) error {

	model, err := t.repo.GetById(ctx, id)
	if err != nil {
		return errors.New("task not found")
	}

	return t.repo.Delete(ctx, model)
}

// GetAllTask implements ITaskService.
func (t *taskService) GetAllTask(ctx context.Context) ([]models.Task, error) {
	return t.repo.GetAll(ctx)
}

// GetById implements ITaskService.
func (t *taskService) GetById(ctx context.Context, id int) (models.Task, error) {
	return t.repo.GetById(ctx, id)
}

// Save implements ITaskService.
func (t *taskService) Save(ctx context.Context, model models.Task) (models.Task, error) {
	return t.repo.Save(ctx, model)
}

// Update implements ITaskService.
func (t *taskService) Update(ctx context.Context, id int, model models.Task) (models.Task, error) {

	oldModel, err := t.repo.GetById(ctx, id)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	oldModel.Title = model.Title
	oldModel.Status = model.Status

	return t.repo.Update(ctx, oldModel)
}

var once sync.Once
var instance *taskService

func NewTaskService(r repositories.ITaskRepository) ITaskService {
	once.Do(func() {
		instance = &taskService{
			repo: r,
		}
	})
	return instance
}
