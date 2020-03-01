package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Usage: goDBMon /path/to/dbmon.xml")
		return
	}

	configMgr := getConfigManager(arguments[1])
	configMgr.readConfig()
	configMgr.showAllDBConfig()
}

func getConfigManager(configFile string) configManager {
	ext := getExtension(configFile)
	switch ext {
	case "xml":
		return &xmlConfigManager{configFile, dbConfig{}}
	case "json":
		return &jsonConfigManager{configFile, dbConfig{}}
	}

	return nil
}

func getExtension(path string) string {
	strArr := strings.Split(path, "/")
	fileName := strArr[len(strArr)-1]
	ext := strings.Split(fileName, ".")[1]
	return ext
}
