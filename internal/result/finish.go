package result

import (
	"github.com/0devs/onecheck/pkg/plugin"
	"os"
)

func FinishAndExit(result *plugin.Result) {
	LogResultWithLogger(result)
	os.Exit(GetExitCode(result))
}
