package model

import (
	"errors"
	"time"
)

var timeLayout = time.DateOnly

type Task struct {
	ID        uint64
	Title     string
	IsDone    bool
	CreatedAt time.Time
	Deadline  time.Time
}

func NewTask(ID uint64, title string, deadlineStr string) (Task, error) {
	isDone := false
	createdAt := time.Now()

	parsedDeadline, err := time.Parse(timeLayout, deadlineStr)

	if err != nil {
		return Task{}, err
	}

	if parsedDeadline.Before(createdAt) {
		return Task{}, errors.New("deadline cannot be past")

	}

	newTask := Task{ID, title, isDone, createdAt, parsedDeadline}

	return newTask, nil
}
