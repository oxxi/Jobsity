package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/oxxi/jobsity/models"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAll(ctx context.Context) ([]models.Task, error)
	GetById(ctx context.Context, id int) (models.Task, error)
	Save(ctx context.Context, model models.Task) (models.Task, error)
	Update(ctx context.Context, model models.Task) (models.Task, error)
	Delete(ctx context.Context, model models.Task) error
}

type taskRepository struct {
	Db *gorm.DB
}

// Delete implements ITaskRepository.
func (t *taskRepository) Delete(ctx context.Context, model models.Task) error {
	if err := t.Db.WithContext(ctx).Delete(&models.Task{}, model.ID).Error; err != nil {
		log.Printf("Error deleting task with id %v: %v", model.ID, err)
		return errors.New("an error occurred while deleting the task")
	}
	return nil
}

// GetAll implements ITaskRepository.
func (t *taskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	if err := t.Db.WithContext(ctx).Order("ID desc").Find(&tasks).Error; err != nil {
		log.Printf("Error getting tasks: %v", err)
		return tasks, errors.New("an error occurred while getting tasks")
	}

	return tasks, nil
}

// GetById implements ITaskRepository.
func (t *taskRepository) GetById(ctx context.Context, id int) (models.Task, error) {
	var task models.Task
	if err := t.Db.WithContext(ctx).First(&task, id).Error; err != nil {
		log.Printf("Error getting task with id %v: %v", id, err)
		return task, fmt.Errorf("an error occurred while getting the task")
	}
	return task, nil
}

// Save implements ITaskRepository.
func (t *taskRepository) Save(ctx context.Context, model models.Task) (models.Task, error) {
	if err := t.Db.WithContext(ctx).Create(&model).Error; err != nil {
		log.Printf("Error saving task: %v", err)
		return model, fmt.Errorf("an error occurred while saving the task")
	}
	return model, nil
}

// Update implements ITaskRepository.
func (t *taskRepository) Update(ctx context.Context, model models.Task) (models.Task, error) {
	if err := t.Db.WithContext(ctx).Save(&model).Error; err != nil {
		log.Printf("Error updating task: %v", err)
		return model, fmt.Errorf("an error occurred while updating the task")
	}
	return model, nil
}

func NewTaskRepository(Db *gorm.DB) ITaskRepository {
	return &taskRepository{Db: Db}
}
