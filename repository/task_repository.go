package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error //全タスクを配列に格納する
	GetTaskById(task *model.Task, UserId uint, TaskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, UserId uint, TaskId uint) error
	DeleteTask(UserId uint, TaskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	// タスク一覧の中から、引数で渡されたuserIdと一致するタスク一覧を取得する
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}
