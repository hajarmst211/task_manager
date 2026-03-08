package main

import (
	"fmt"
	"taskManager/service"
	"taskManager/ui"
)

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
	ui.MenuHandler(choice, personnalManager)
}
