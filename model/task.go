package task

import (
	"errors"
	"taskManager/helpers"
	"log"
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

func NewTask(title string,isDone bool, createdAt time.Time, deadline time.Time) (Task, error) {

	if deadline.Before(createdAt) {
		return Task{}, errors.New("deadline cannot be past")

	}

	newTask := Task{
		 Title : title,
		 IsDone:isDone,
		 CreatedAt: createdAt,
		 Deadline:deadline,
		}

	return newTask, nil
}

func SaveTask(task Task) error{
	db := database.OpenDB()

	query := "INSERT INTO tasks (title, createdAT, deadline) VALUES (?,?,?);"

	result, err := db.Exec(query, task.Title, task.CreatedAt, task.Deadline)
	if err != nil{
		log.Fatal(err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return nil
	}
	task.ID = uint64(id)

	database.CloseDB(db)
	return nil
}

func ListTasks() []Task{
	db := database.OpenDB()
	tasksList := [] Task{};

	query := "SELECT * FROM tasks ORDER BY deadline;"

	rows, err := db.Query(query)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title string
		var isDone bool
		var createdAt time.Time
		var deadline time.Time

		err := rows.Scan(&id, &title, &isDone, &createdAt, &deadline)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue 
		}
		newTask, err := NewTask(title, isDone, createdAt, deadline)
		tasksList = append(tasksList, newTask)
	}

	return tasksList	
}

