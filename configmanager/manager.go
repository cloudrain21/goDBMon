package configmanager

type ConfigManager interface {
	ReadConfig() error
	ShowConfig()
}