/*
 * ISBHI2Shell
 *
 * Copyright (c) 2026 ISBHI
 *
 * File: navigation.go
 * Description: Implements navigation commands.
 */

package commands

import (
	"fmt"
	"isbhi2shell/config"
	"os"
	"path/filepath"
	"strings"
)

func Ls(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("\n" + err.Error())
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry.Name() + "/")
		} else {
			fmt.Println(entry.Name())
		}
	}
}

func Cd(path string) {
	if path == "~" {
		path = config.UserHomeDir()
	} else if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, `~\`) {
		path = filepath.Join(config.UserHomeDir(), path[2:])
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Println("\n" + err.Error())
	}
}

func PromptWD() string {
	rel, err := filepath.Rel(config.UserHomeDir(), Pwd())

	if err != nil {
		fmt.Println("\n" + err.Error())
		return ""
	}

	return filepath.Join("~", rel)
}

func Pwd() string {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("\n" + err.Error())
		return ""
	}

	return dir
}
