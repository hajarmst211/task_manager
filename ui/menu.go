package ui

import (
	"fmt"
)


func MainMenu()  uint8{
	var main_menu_choice uint8

	fmt.Println(`
	Select an action:
	(1) Add task
	(2) Print all tasks
	(3) Edit a deadline
	(4) Clear tasks
	`)

	fmt.Scanln(&main_menu_choice)
	return main_menu_choice
}

func AddTaskMenu() (string, string){
	var title string
	var deadline string

	fmt.Println("Enter the title of the task:")
	fmt.Scanln(&title)

	fmt.Println("Enter the deadline:")
	fmt.Scanln(&deadline)

	return title, deadline
}
