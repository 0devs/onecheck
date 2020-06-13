package command

import (
	filePlugin "github.com/0devs/onecheck/plugin/file"
)

type CommandParams struct {
	Path   string
	Exists bool
}

var commandParams CommandParams

func convertCommandParamsToConfig(params *CommandParams) filePlugin.Config {
	config := filePlugin.Config{
		Path:   params.Path,
		Exists: params.Exists,
	}

	return config
}
