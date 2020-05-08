package main

import (
	cmgr "github.com/cloudrain21/goDBMon/configmanager"
	"github.com/cloudrain21/goDBMon/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func init() {
	logger.Configure(logger.NewFileLogger("dbmon.log"))
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		log.Println("Usage: goDBMon /path/to/dbmon.xml")
		return
	}

	configFile := arguments[1]

	configManager := CreateConfigManager(configFile)
	err := configManager.Mgr.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	configManager.Mgr.ShowConfig()

}

func CreateConfigManager(configFile string) *cmgr.ConfigManager {
	l := strings.Split(configFile, ".")
	ext := l[len(l)-1]

	switch ext {
	case "xml":
		log.Println("xml file extension")
		return cmgr.NewConfigManager(cmgr.NewXmlConfigManager(configFile))
	case "json":
		log.Println("json file extension")
		return cmgr.NewConfigManager(cmgr.NewJsonConfigManager(configFile))
	default:
		log.Println("default file extension")
		return cmgr.NewConfigManager(cmgr.NewXmlConfigManager(configFile))
	}
}
