package configmanager

import (
	log "github.com/sirupsen/logrus"
)

type JsonManager struct {
	configFile string
}

func (j *JsonManager) ReadConfig() error {
	log.Println("read json config...")
	return nil
}

func (j JsonManager) ShowConfig() {
	log.Println("show json config...")
}
