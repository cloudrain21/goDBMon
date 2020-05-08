package configmanager

import "errors"

var (
	InvalidExtension = errors.New("invalid config file extension")
	InvalidSyntax    = errors.New("invalid config syntax")
)

type InterConfigManager interface {
	ReadConfig() error
	ShowConfig()
}

type ConfigManager struct {
	Mgr InterConfigManager
}

func NewConfigManager(mgrinter InterConfigManager) *ConfigManager {
	cmgr := new(ConfigManager)

	cmgr.Mgr = mgrinter
	return cmgr
}
