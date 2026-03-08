package main

import (
	"fmt"
	"taskManager/model"
	"taskManager/service"
	"taskManager/ui"
	"time"
)

var timeLayout = time.DateOnly

func printTask(task model.Task) string {
	status := "not done"
	if task.IsDone {
		status = "done"
	}
	return fmt.Sprintf("ID: %d, Title: %s, status: %s, deadline: %s", task.ID, task.Title, status, task.Deadline.Format(timeLayout))

}

func printTasks(taskManager *service.TaskManager) {
	taskManager.Mu.RLock()
	TasksHash := taskManager.TasksMap
	if len(TasksHash) == 0 {
		fmt.Println("No task available")
		return
	}

	fmt.Println("----Current tasks are:")
	for _, task := range TasksHash {
		fmt.Println(printTask(task))
	}
	fmt.Println("----done printing")
	taskManager.Mu.Unlock()
}

func PritnTodayDeadline(taskManager *service.TaskManager) {
	tasks := taskManager.GetTodayDeadline()
	for _, task := range tasks {
		printTask(task)
	}
}

func main() {
	personnalManager := service.NewTaskManager()

	addingTaskError := personnalManager.AddTask("finish manager", "2026-05-02")
	if addingTaskError != nil {
		fmt.Print(addingTaskError)
		return
	}

	fmt.Printf("The personnal manager tasksmap is: %v\n", personnalManager.TasksMap)

	personnalManager.UpdateDeadlineTitle("finish manager", "2027-04-05")

	choice := ui.MainMenu()
	fmt.Println("Le choix est", choice)

}
