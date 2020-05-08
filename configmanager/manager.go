package configmanager

import "errors"

var (
	InvalidExtension = errors.New("invalid config file extension")
	InvalidSyntax    = errors.New("invalid config syntax")
)

type ConfigManager interface {
	ReadConfig() error
	ShowConfig()
}