package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var globalIDCounter int64

type Task struct {
	ID        int64
	Title     string
	isDone    bool
	createdAt time.Time
	Deadline  time.Time
}

func createTask(title string, deadlineStr string) Task {
	newID := atomic.AddInt64(&globalIDCounter, 1)
	isDone := false
	createdAt := time.Now()

	layout := time.DateOnly
	parsedDeadline, err := time.Parse(layout, deadlineStr)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		return Task{}
	}

	if parsedDeadline.Before(createdAt) {
		fmt.Println("The deadline is before than today, change that")
		return Task{}
	}

	newTask := Task{newID, title, isDone, createdAt, parsedDeadline}

	return newTask
}

func addTask(task Task, TasksHash map[int64]Task) {
	if _, isPresent := TasksHash[task.ID]; isPresent {
		fmt.Println("A task with this id exists already, try again!")
	} else {
		TasksHash[task.ID] = task
	}
}

func 

func main() {
	var TasksHash = make(map[int64]Task)
	newTask := createTask("finish this planner", "2023-10-27")

	if newTask == (Task{}) {
		fmt.Println("The task is not created it is empty")
		return
	}

	addTask(newTask, TasksHash)

	fmt.Printf("The task id is: %d \n The title is: %s \n The deadline is: %v \n It was created at: %v \n",
		newTask.ID, newTask.Title, newTask.Deadline.Format("2025-01-14"), newTask.createdAt.Format("2025-01-14"))

	fmt.Println("the hashmap is:", TasksHash)
}
