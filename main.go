package main

import (
	"taskManager/model"
	"time"
	"fmt"
)
func main() {
	layout := "2006-01-02"
	deadline,_ := time.Parse(layout, "2025-05-01")
	createdAT, _:= time.Parse(layout, "2025-05-01")
	newTask, err := task.NewTask("finish planner", false, createdAT, deadline)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(newTask)
}
