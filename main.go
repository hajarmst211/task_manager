package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"taskManager/helpers"
)


func main() {
	clearScreen()
	todayTasks, err := helpers.TodayTasks()
	if err != nil {
		fmt.Printf("Error occured: %s\n", err)
	}

	fmt.Println("======= DAILY TASK MANAGER =======")
	fmt.Println("======= Tasks that end TODAY: =====")
	helpers.PrintNTasks(todayTasks)

	for {
		err := helpers.MainMenu()
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			continue
		}
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
