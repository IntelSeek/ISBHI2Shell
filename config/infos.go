package config

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

const Version = "v0.1.1"

func OS() string {
	return runtime.GOOS
}

func CPUs() int {
	return runtime.NumCPU()
}

func Arch() string {
	return runtime.GOARCH
}

func Username() string {
	currentUser, err := user.Current()

	if err != nil {
		fmt.Println("\nCan't get username:", err)
		return ""
	}

	username := currentUser.Username

	if runtime.GOOS == "windows" {
		if i := strings.LastIndex(username, `\`); i >= 0 {
			username = username[i+1:]
		}
	}

	return username
}

func HostName() string {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("\nCan't get hostname:", err)
		return ""
	}

	return hostname
}

var Reference = Username() + " - ISBHI"

func UserHomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return dir
}

const (
	Name      = "ISBHI2Shell"
	Developer = "ISBHI - IntelSeek Bureau of Hacking and Intelligence"
)
