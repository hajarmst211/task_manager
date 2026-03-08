package ui

import "fmt"

func MainMenu()  string{
	var main_menu_choice string
	fmt.Print("Enter the number of the action you want to perform:\n   (1): Add task\n	(2):Print a task\n    (3):Print all tasks\n   (4):Edit a deadline\n    (5):Clear tasks")

	fmt.Scanln(&main_menu_choice)
	return main_menu_choice
}

