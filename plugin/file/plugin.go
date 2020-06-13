package filePlugin

import (
	"fmt"
	"github.com/0devs/onecheck/pkg/plugin"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type FilePlugin struct {
	Name    string
	Version string
}

func (p *FilePlugin) ValidateConfig(config *Config) error {
	return validateConfig(config)
}

func (p *FilePlugin) Run(config *Config) (error, plugin.Result) {
	logger := zap.L()
	logger.Debug("run", zap.Any("config", config))

	result := plugin.Result{}

	fs := afero.NewOsFs()

	exists, err := afero.Exists(fs, config.Path)

	if err != nil {
		return err, result
	}

	result.SetStatusMessage(
		fmt.Sprintf("path=%s, shouldExists=%t, exists=%t", config.Path, config.Exists, exists),
	)

	if exists == config.Exists {
		result.SetStatusOk()
	} else {
		result.SetStatusWarn()
	}

	return nil, result
}

func (p *FilePlugin) ValidateConfigFromMap(configMap interface{}) error {
	logger := zap.L()

	logger.Debug("configMap", zap.Any("configMap", configMap))

	return nil
}

//func (p *FilePlugin) RunFromMap(configMap interface{}) {
//
//}

var Plugin = FilePlugin{
	Name:    "file",
	Version: "0.0.0",
}
