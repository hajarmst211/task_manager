package ui

import(
	"taskManager/service"
	"fmt"
)

func addTaskHandler(taskManager *service.TaskManager){
	title, deadline := AddTaskMenu()
	taskManager.AddTask(title, deadline)
}


func updateDeadlineHandler(taskManager *service.TaskManager){
	var title string
	var id uint64
	var deadline string
	var choice uint8
	var updateError error

	fmt.Println(`
				Do you want to select the task by:
				(1): ID
				(2): title
				`)
	fmt.Scanln(choice)

	if choice == 1{

		fmt.Println("Enter the ID:")
		fmt.Scanln(&id)

		fmt.Println("Enter the new deadline (format: yyyy-mm-dd):")
		fmt.Scanln(&deadline)

		updateError = taskManager.UpdateDeadlineID(id,deadline)

	}else{
		fmt.Println("Enter the title:")
		fmt.Scanln(&title)

		fmt.Println("Enter the new deadline (format: yyyy-mm-dd):")
		fmt.Scanln(deadline)

		updateError = taskManager.UpdateDeadlineTitle(title,deadline)
	}

	if updateError == nil{
			fmt.Println("Update succeful")
			return
		}
	fmt.Println("The update wasn't succeful")
		return

}

func MenuHandler(choice uint8, taskManager *service.TaskManager){
	switch choice{
	case 1: addTaskHandler(taskManager)
	case 2: PrintTasks(taskManager)
	case 3:	updateDeadlineHandler(taskManager)
	case 4: taskManager.ClearTasks() 

	}
}