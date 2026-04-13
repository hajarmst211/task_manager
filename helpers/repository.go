package helpers

import (
	"encoding/json"
	"os"
	"taskManager/model"
	"time"
)

const jsonPath =  "../tasks.json"

func SaveTask(task model.Task) error{
	file, err := os.Create(jsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(task)
}


func LoadTasks() ([]model.Task, error){
	file, err := os.Open(jsonPath)
	if err != nil{
		return  []model.Task{} ,nil
	}
	defer file.Close()
	var tasks  []model.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	return tasks, err
}

func CreateTask(title string, deadline time.Time, details string) (model.Task, error){
	tasks, err := LoadTasks()
	if err != nil{
		return model.Task{}, err
	}

	newID := 0
	if len(tasks) > 0{
		maxID := 0
		for _, task := range tasks {
			if task.ID > maxID{
				maxID = task.ID
			}
		} 
		newID = maxID +1
	}

	newTask := model.Task{
		ID: newID,
		Title: title,
		Deadline: deadline,
		Details: details,
	}
	
	return newTask, nil
}

func MarkDone(ID int) error{
	tasks, err := LoadTasks()
	if err != nil{
		return err
	}
	isFound := false

	for _, task := range tasks{
		if task.ID == ID{
			task.Done = true
			isFound = true
			break
		}
	}

	if isFound{
		
	}
	
}