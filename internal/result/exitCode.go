package result

import "github.com/0devs/onecheck/pkg/plugin"

func GetExitCode(result *plugin.Result) int {
	switch (result.Status) {
	case "ok":
		return 0

	case "warn":
		return 1

	case "error":
		return 2

	default:
		return 254
	}
}