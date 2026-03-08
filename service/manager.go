package service

import (
	"fmt"
	"taskManager/model"
	"sync/atomic"
	"sync"
	"errors"
	"time"
)

var timeLayout = time.DateOnly
var nextID uint64


type TaskManager struct {
	Mu       sync.RWMutex
	TasksMap map[uint64]model.Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		TasksMap: make(map[uint64]model.Task),
	}
}


func (taskManager *TaskManager) AddTask(title string, deadlineStr string) error {
	taskManager.Mu.Lock()
	defer taskManager.Mu.Unlock()
	
	newID := atomic.AddUint64(&nextID, 1)

	taskToAdd, taskError := model.NewTask(uint64(newID), title, deadlineStr)
	if taskError != nil {
		return taskError
	}

	taskManager.TasksMap[newID] = taskToAdd
	return nil
}

func (taskManager *TaskManager) DeleteTask(ID uint64) {
	taskManager.Mu.Lock()
	defer taskManager.Mu.Unlock()
	delete(taskManager.TasksMap, ID)
}

func (taskManager *TaskManager) GetTask(ID uint64) model.Task {
	taskManager.Mu.RLock()
	defer taskManager.Mu.RUnlock()
	return taskManager.TasksMap[ID]
}

func (taskManager *TaskManager) UpdateDeadlineID(ID uint64, deadline string) error {
	taskManager.Mu.Lock()
	defer taskManager.Mu.Unlock()
	
	timeDeadline, error := time.Parse(timeLayout, deadline)
	if error != nil {
		return errors.New("invalid date format")
	}
	task, exist := taskManager.TasksMap[ID]
	if !exist {
		return fmt.Errorf("No task exists with this ID:%d", ID)
	}
	task.Deadline = timeDeadline
	taskManager.TasksMap[ID] = task
	return nil
}

func (taskManager *TaskManager) UpdateDeadlineTitle(title string, deadline string) error {
	taskManager.Mu.Lock()
	defer taskManager.Mu.Unlock()

	for _, task := range taskManager.TasksMap {
		if task.Title == title {
			ID := task.ID
			// Removed internal Lock/Unlock from UpdateDeadlineID call logic
			// By setting the values directly here, we avoid the recursive deadlock
			timeDeadline, err := time.Parse(timeLayout, deadline)
			if err != nil {
				return errors.New("invalid date format")
			}
			task.Deadline = timeDeadline
			taskManager.TasksMap[ID] = task
			return nil
		}
	}
	return fmt.Errorf("No task exists witht this title")
}

func (taskManager *TaskManager) ClearTasks(){
	taskManager.Mu.Lock()
	defer taskManager.Mu.Unlock()
	taskManager.TasksMap = make(map[uint64]model.Task)
}

func(taskManager *TaskManager) GetTodayDeadline() []model.Task{
	taskManager.Mu.RLock()
	defer taskManager.Mu.RUnlock()
	
	var tasksList []model.Task
	today := time.Now().Format(time.DateOnly)

	for _, task := range taskManager.TasksMap{
		taskDate := task.Deadline.Format(time.DateOnly)
		if taskDate== today{
			tasksList = append(tasksList, task)
		}
	}
	return tasksList
}
