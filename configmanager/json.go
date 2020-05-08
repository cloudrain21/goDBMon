package configmanager

import (
	log "github.com/sirupsen/logrus"
)

type JsonConfigManager struct {
	configFile string
}

func (j *JsonConfigManager) ReadConfig() error {
	log.Println("read json config...")
	return nil
}

func (j JsonConfigManager) ShowConfig() {
	log.Println("show json config...")
}

func NewJsonConfigManager(configFile string) ConfigManager {
	return &JsonConfigManager{configFile}
}