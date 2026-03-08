package ui


import (
	"fmt"
	"taskManager/model"
	"taskManager/service"
	"time"
)

var timeLayout = time.DateOnly

func PrintTask(task model.Task) string {
	status := "not done"
	if task.IsDone {
		status = "done"
	}
	return fmt.Sprintf("ID: %d, Title: %s, status: %s, deadline: %s", task.ID, task.Title, status, task.Deadline.Format(timeLayout))

}

func PrintTasks(taskManager *service.TaskManager) {
	taskManager.Mu.RLock()
	TasksHash := taskManager.TasksMap
	if len(TasksHash) == 0 {
		fmt.Println("No task available")
		return
	}

	fmt.Println("----Current tasks are:")
	for _, task := range TasksHash {
		fmt.Println(PrintTask(task))
	}
	fmt.Println("----done printing")
	taskManager.Mu.Unlock()
}

func PritnTodayDeadline(taskManager *service.TaskManager) {
	tasks := taskManager.GetTodayDeadline()
	for _, task := range tasks {
		PrintTask(task)
	}
}