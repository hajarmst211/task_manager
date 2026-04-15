package helpers

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func MainMenu() error {

	fmt.Print(`MENU:

    1. List Tasks

    2. Add Task

    3. Mark Task as Done

    4. Delete Task

    5. Quit
`)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var choice int
	_, err := fmt.Sscanf(input, "%d", &choice)
	if err != nil {
		return fmt.Errorf("please enter a valid number")
	}

	if choice > 5 || choice < 1 {
		return fmt.Errorf("This number is out of the menu range")
	}

	switch choice {
	case 1:
		PrintTasks()
	case 2:
		err = AddTaskMenu()
	case 3:
		err = MarkDoneMenu()
	case 4:
		err = DeleteMenu()
	case 5:
		fmt.Println("Exiting...")
		os.Exit(0)
		return nil
	}

	if err != nil {
		return fmt.Errorf("\nTask not completed: %w!\n Exiting...", err)
	}

	return nil
}

func DeleteMenu() error {
	fmt.Println("Enter the ID of the task to delete:")

	var ID int
	fmt.Scanln(&ID)

	err := DeleteTask(ID)

	if err != nil {
		return err
	}

	fmt.Println("Task deleted succefully!")
	return nil
}

func MarkDoneMenu() error {
	fmt.Println("Enter the ID of the task to mark done:")

	var ID int
	fmt.Scanln(&ID)

	_, err := LoadTask(ID)

	if err != nil {
		return err
	}

	err = MarkDone(ID)

	if err != nil {
		return err
	}

	fmt.Println("Task is done... You are doing a good job Hajora!")
	return nil
}

func AddTaskMenu() error {
	fmt.Println("Enter informations about the task to create:")

	scanner := bufio.NewScanner(os.Stdin)
	var title string

	for {
		fmt.Println("Enter the title:")
		scanner.Scan()
		title = scanner.Text()
		if title != "" {
			break
		}
		fmt.Println("Title cannot be empty!")
	}

	var parsedDeadline time.Time
	for {
		today := time.Now().Truncate(24 * time.Hour)
		fmt.Printf("Today is: %s. Enter the deadline:", today)
		scanner.Scan()
		deadline := scanner.Text()

		parsedDeadline, err := DateParser(deadline)
		if err != nil {
			fmt.Println("Invalid format! Please use YYYY-MM-DD.")
			continue
		}

		if parsedDeadline.Before(today) {
			fmt.Printf("The deadline (%s) is in the past! Today is %s.\n",
				parsedDeadline.Format("2006-01-02"),
				today.Format("2006-01-02"))
			continue
		}
		break
	}

	fmt.Println("Do you have other details about the task?(optional):")
	scanner.Scan()
	details := scanner.Text()

	newTask, err := CreateTask(title, parsedDeadline, details)
	if err != nil {
		return err
	}

	err = AddTask(newTask)

	if err != nil {
		return fmt.Errorf("Error adding the task:%w", err)
	}
	fmt.Printf("task was added succefully! \n make sure YOU DO IT!\n")
	return nil
}
