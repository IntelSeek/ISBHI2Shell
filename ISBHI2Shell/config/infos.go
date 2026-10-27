package config

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

const Version = "v0.1.0"

var Username = func() string {
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
}()

var Reference = Username + " - ISBHI"

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
