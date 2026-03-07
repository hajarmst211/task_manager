package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)


type Task struct {
	ID        uint64
	Title     string
	isDone    bool
	createdAt time.Time
	Deadline  time.Time
}

type TaskManager struct {
	mu sync.Mutex
	tasksMap map[uint64]Task 
	nextID uint64 
}

func newTaskManager() *TaskManager{
	return &TaskManager{
		tasksMap: make(map[uint64]Task),
		nextID: 1,
	}
}

func newTask(ID uint64, title string, deadlineStr string) (Task, error) {
	isDone := false
	createdAt := time.Now()

	layout := time.DateOnly
	parsedDeadline, err := time.Parse(layout, deadlineStr)

	if err != nil {
		return Task{}, err
	}

	if parsedDeadline.Before(createdAt) {
		return Task{}, errors.New("deadline cannot be past")

	}

	newTask := Task{ID, title, isDone, createdAt, parsedDeadline}

	return newTask, nil
}


func (taskManager *TaskManager) addTask(title string, deadlineStr string) error {
	TasksHash := taskManager.tasksMap
	newID := taskManager.nextID
	
	taskToAdd, taskError := newTask(uint64(newID), title, deadlineStr)
	if taskError != nil{
		return taskError
	}

	TasksHash[newID] = taskToAdd
	taskManager.nextID += 1
	return nil
}

func (taskManager *TaskManager)deleteTask(ID int64) {
	TasksHash := taskManager.tasksMap
	delete(TasksHash, ID)
}

func (taskManager *TaskManager)printTasks() {
	TasksHash := taskManager.tasksMap
	for key, value := range TasksHash {
		fmt.Printf("The value: %d is: \n	%v", key, value)
	}
}

func main() {
	personnalManager := newTaskManager()
	finishPlanerTask, error := newTask("finish this planner", "2028-10-27")

	if error == nil {
		fmt.Println("The task is not created it is empty")
		return
	}

	personnalManager.addTask(finishPlanerTask)

}
