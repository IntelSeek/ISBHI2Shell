/*
 * ISBHI2Shell
 *
 * Copyright (c) 2026 ISBHI
 *
 * File: basic.go
 * Description: Implements basic commands such as help and whoami.
 */

package commands

import "fmt"

func Help(args []string) {
	if len(args) == 0 {
		fmt.Println(`
help    - shows a list of all commands
version - outputs current ISBHI2Shell version
about   - shows info about ISBHI2Shell
whoami  - outputs current user
pwd     - shows current working directory
cd      - changes current working directory
ls      - lists files and directories
sysinfo - shows info about system and ISBHI2Shell
clear   - clears terminal
exit    - exits ISBHI2Shell`)
		return
	}

	for _, arg := range args {
		switch arg {
		case "help":
			fmt.Println(`
help - shows a list of all commands
Usage: help <(opt) cmd ...>`)

		case "version":
			fmt.Println(`
version - outputs current ISBHI2Shell version
Usage: version`)

		case "about":
			fmt.Println(`
about - shows info about ISBHI2Shell
Usage: about`)

		case "whoami":
			fmt.Println(`
whoami - outputs current user
Usage: whoami`)

		case "pwd":
			fmt.Println(`
pwd - shows current working directory
Usage: pwd`)

		case "cd":
			fmt.Println(`
cd - changes current working directory
Usage: cd <(opt) path>
Examples:
  cd
  cd ~
  cd ..
  cd Documents`)

		case "ls":
			fmt.Println(`
ls - lists files and directories
Usage: ls <(opt) path>
Examples:
  ls
  ls ..
  ls Documents`)

		case "sysinfo":
			fmt.Println(`
sysinfo - shows info about system and ISBHI2Shell
Usage: sysinfo`)

		case "clear":
			fmt.Println(`
clear - clears terminal
Usage: clear`)

		case "exit":
			fmt.Println(`
exit - exits ISBHI2Shell
Usage: exit`)

		default:
			fmt.Println("\nNo help available for command:", arg)
		}
	}
}

func Whoami(username string) {
	fmt.Println("\n" + username)
}

func Clear() {
	fmt.Print("\033[H\033[2J\033[3J")
}
