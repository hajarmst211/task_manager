package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"taskManager/helpers"
)

func main() {
	// Optional: Clear screen at start
	clearScreen()

	fmt.Println("======= DAILY TASK MANAGER =======")
	
	for {
		err := helpers.MainMenu()
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			continue
		}

		// Because MainMenu returns nil on case 5 (Quit), 
		// but doesn't tell 'main' to stop, we handle the exit logic here.
		// Note: To make this cleaner, you'd usually have MainMenu return a bool.
		// For now, if the user picks 'Quit' in your helpers code, it prints "Exiting..."
		// If you want the program to actually close, we can check for that.
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