package main

import (
	"fmt"
	"log"
	"taskManager/helpers"
	"time"
)

func main() {
	// 1. Create a new task
	fmt.Println("--- Creating a Task ---")
	deadline := time.Now().AddDate(0, 0, 7) // 7 days from now
	task1, err := helpers.CreateTask("Learn Go", deadline, "Study slices and maps")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Add the task to the JSON file
	err = helpers.AddTask(task1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added Task: %s (ID: %d)\n", task1.Title, task1.ID)

	// 3. Add a second task to test ID incrementing
	task2, _ := helpers.CreateTask("Buy Milk", time.Now(), "Get full cream milk")
	helpers.AddTask(task2)
	fmt.Printf("Added Task: %s (ID: %d)\n", task2.Title, task2.ID)

	// 4. Load and Print all tasks
	fmt.Println("\n--- Current Task List ---")
	printTasks()

	// 5. Mark Task 0 as Done
	fmt.Println("\n--- Marking Task 0 as Done ---")
	err = helpers.MarkDone(0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	printTasks()

	// 6. Delete Task 1
	fmt.Println("\n--- Deleting Task 1 ---")
	err = helpers.DeleteTask(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	printTasks()
}

func printTasks() {
	tasks, err := helpers.LoadTasks()
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "X"
		}
		fmt.Printf("[%s] ID: %d | Title: %s | Deadline: %s\n", status, t.ID, t.Title, t.Deadline.Format("2006-01-02"))
	}
}