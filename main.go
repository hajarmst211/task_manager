package main

import (
//	"fmt"
	"taskManager/helpers"
//	"taskManager/model"
//	"time"
)
func main() {
/*	deadline,_ := time.Parse(layout, "2025-05-01")
	createdAT, _:= time.Parse(layout, "2025-05-01")
	newTask, err := task.NewTask("finish planner", false, createdAT, deadline)
	if err != nil{
		fmt.Println(err)
	}

	database.Trydb()
	task_err := SaveTask(newTask)

	fmt.Println("this is the new task error:" ,task_err)
	*/
database.Trydb()	
}
