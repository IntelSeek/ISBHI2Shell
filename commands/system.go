package commands

import (
	"fmt"
	"isbhi2shell/config"
)

func SysInfo() {
	fmt.Printf("\n%-12s %s\n", "OS", config.OS())
	fmt.Printf("%-12s %s\n", "ARCH", config.Arch())
	fmt.Printf("%-12s %s\n", "ISBHI2Shell", config.Version)
	fmt.Printf("%-12s %d\n", "CPUs", config.CPUs())
	fmt.Printf("%-12s %s\n", "User", config.Username())
	fmt.Printf("%-12s %s\n", "Home", config.UserHomeDir())
}
