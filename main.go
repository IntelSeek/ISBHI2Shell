package main

import (
	"bufio"
	"fmt"
	"isbhi2shell/commands"
	"isbhi2shell/config"
	"os"
	"strings"
)

func main() {
	commands.Cd("~")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ISBHI2Shell", config.Version+"\nType 'help' to list commands.")

	for {
		fmt.Print("\n╭──[ " + config.Username() + "@" + config.HostName() + " ]-(" + commands.PromptWD() + ")\n╰─$ ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("\nInput error:", err)
			}
			return
		}

		parts := strings.Fields(scanner.Text())

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "help":
			commands.Help(args)

		case "version":
			fmt.Println("\n"+config.Name, config.Version)

		case "about":
			fmt.Println("\n" + config.Name + " " + config.Version +
				"\nDeveloped by " + config.Developer +
				"\n" + config.Reference)

		case "whoami":
			commands.Whoami(config.Username())

		case "pwd":
			fmt.Println("\n" + commands.Pwd())

		case "cd":
			if len(args) == 0 {
				commands.Cd("~")
			} else if len(args) == 1 {
				commands.Cd(args[0])
			} else {
				fmt.Println("\nToo many arguments for cd.")
			}

		case "ls":
			if len(args) == 0 {
				commands.Ls(".")
			} else if len(args) == 1 {
				commands.Ls(args[0])
			} else {
				fmt.Println("\nToo many arguments for ls.")
			}

		case "sysinfo":
			commands.SysInfo()

		case "clear":
			commands.Clear()

		case "exit":
			fmt.Println("\nBye.")
			return

		default:
			fmt.Println("\nUnknown command.")
		}
	}
}
