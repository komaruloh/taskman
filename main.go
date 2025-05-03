/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/komaruloh/taskman/cmd"
	"github.com/komaruloh/taskman/internal/pkg/config"
	"github.com/komaruloh/taskman/internal/pkg/db"
	"github.com/komaruloh/taskman/internal/pkg/utils"
	gap "github.com/muesli/go-app-paths"
)

func main() {
	initApp()
	cmd.Execute()
}

func initApp() {
	scope := gap.NewScope(gap.User, "taskman")

	_, err := config.NewConfig(scope)
	if err != nil {
		panic(err)
	}

	dataDir, err := setupDataPath(scope)
	if err != nil {
		panic(err)
	}

	_, err = setupConfigPath(scope)
	if err != nil {
		panic(err)
	}

	if err := db.SetupDb(dataDir); err != nil {
		panic(err)
	}

}

// create app home directory if it doesn't exist
func setupDataPath(scope *gap.Scope) (string, error) {
	dataDirs, err := scope.DataDirs()
	if err != nil {
		return "", err
	}

	var dataDir string
	if len(dataDirs) > 0 {
		dataDir = dataDirs[0]
	} else {
		dataDir, err = os.UserHomeDir()
		if err != nil {
			return "", err
		}
	}

	if err := utils.InitDir(dataDir); err != nil {
		return "", err
	}
	return dataDir, nil
}

// create app config directory if it doesn't exist
func setupConfigPath(scope *gap.Scope) (string, error) {
	configDirs, err := scope.ConfigDirs()
	if err != nil {
		return "", err
	}

	var configDir string
	if len(configDirs) > 0 {
		configDir = configDirs[0]
	} else {
		configDir, err = os.UserHomeDir()
		if err != nil {
			return "", err
		}
	}

	if err := utils.InitDir(configDir); err != nil {
		return "", err
	}

	return configDir, nil
}
