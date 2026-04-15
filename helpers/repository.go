package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"taskManager/model"
	"text/tabwriter"
	"time"
)

const jsonPath = "tasks.json"

func OverwriteTasks(tasks []model.Task) error {
	file, err := os.Create(jsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(tasks)
}

func LoadTasks() ([]model.Task, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return []model.Task{}, nil
	}
	defer file.Close()
	var tasks []model.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	return tasks, err
}

func LoadTask(ID int) (model.Task, error){
	tasks, err := LoadTasks()
	if err != nil{
		return model.Task{},nil
	}

	for _, task := range tasks{
		if task.ID == ID{
			return task, nil
		}
	}

	return model.Task{}, fmt.Errorf("No task with that ID was found!")
}


func AddTask(newTask model.Task) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, newTask)
	err = OverwriteTasks(tasks)
	return err
}

func CreateTask(title string, deadline time.Time, details string) (model.Task, error) {
	now := time.Now().Format(TimeLayout)
	createdAt, err := DateParser(now)

	if err != nil {
		panic(err)
	}

	if deadline.Before(createdAt) {
		return model.Task{}, fmt.Errorf("The deadline is past today")
	}

	tasks, err := LoadTasks()
	if err != nil {
		return model.Task{}, err
	}

	newID := 0
	if len(tasks) > 0 {
		maxID := 0
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		newID = maxID + 1
	}

	newTask := model.Task{
		ID:        newID,
		Title:     title,
		CreatedAt: createdAt,
		Deadline:  deadline,
		Details:   details,
	}

	return newTask, nil
}

func MarkDone(ID int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	isFound := false

	for i := range tasks {
		if tasks[i].ID == ID {
			tasks[i].Done = true
			isFound = true
			break
		}
	}

	if !isFound {
		return fmt.Errorf("task with ID %d not found", ID)
	}

	OverwriteTasks(tasks)
	return nil
}

func DeleteTask(ID int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return nil
	}

	index := -1
	for i := range tasks {
		if tasks[i].ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with the ID %d not found", ID)
	}

	tasks1 := tasks[:index]
	tasks2 := tasks[index+1:]

	tasks = append(tasks1, tasks2...)

	return OverwriteTasks(tasks)
}

func PrintTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found. Your todo list is empty!")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Header
	fmt.Fprintln(w, "STATUS\tID\tTITLE\tDEADLINE\tCREATED")
	fmt.Fprintln(w, "------\t--\t-----\t--------\t-------")

	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[X]"
		}

		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\n",
			status,
			t.ID,
			t.Title,
			t.Deadline.Format("2006-01-02"),
			t.CreatedAt.Format("2006-01-02"),
		)
	}
	w.Flush()
}
