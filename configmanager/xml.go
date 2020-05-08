package configmanager

import (
	log "github.com/sirupsen/logrus"
)

type XmlConfigManager struct {
	configFile string
}

func (x *XmlConfigManager) ReadConfig() error {
	log.Printf("read config from xml config file : %s\n", x.configFile)
	return nil
}

func (x XmlConfigManager) ShowConfig() {
	log.Printf("show configuration : %s", x.configFile)
}

func NewXmlConfigManager(configFile string) ConfigManager {
	return &XmlConfigManager{configFile}
}
