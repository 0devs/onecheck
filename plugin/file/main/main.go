package main

import (
	filePlugin "github.com/0devs/onecheck/plugin/file"
	"github.com/0devs/onecheck/plugin/file/command"
)

var Name string = filePlugin.Plugin.Name
var Version string = filePlugin.Plugin.Version

var Init command.InitFn = command.Init