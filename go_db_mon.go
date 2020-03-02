package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Usage: goDBMon /path/to/dbmon.xml")
		return
	}

	configMgr, err := getConfigManager(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = configMgr.readConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	configMgr.showAllDBConfig()
}

var errInvalidExtension = errors.New("invalid config file extension")

func getConfigManager(configFile string) (configManager, error) {
	ext := getExtension(configFile)
	switch ext {
	case "xml":
		return &xmlConfigManager{configFile, map[string]DBConfigMap{}}, nil
	case "json":
		return &xmlConfigManager{configFile, map[string]DBConfigMap{}}, nil
	default:
		return nil, errInvalidExtension
	}
}

func getExtension(path string) string {
	fileName := filepath.Base(path)
	ext := strings.Split(fileName, ".")[1]
	return ext
}
