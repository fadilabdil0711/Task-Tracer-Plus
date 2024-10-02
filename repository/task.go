package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	result := t.db.Model(&model.Task{}).Where("id = ?", task.ID).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *taskRepository) Delete(id int) error {
	result := t.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var tasks []model.Task
	result := t.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var taskCategories []model.TaskCategory
	result := t.db.Table("tasks").Select("tasks.*, categories.name as category").
		Joins("LEFT JOIN categories ON tasks.category_id = categories.id").
		Where("tasks.id = ?", id).
		Find(&taskCategories)
	if result.Error != nil {
		return nil, result.Error
	}
	return taskCategories, nil
}
