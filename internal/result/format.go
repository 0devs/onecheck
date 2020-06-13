package result

import (
	"fmt"
	"github.com/0devs/onecheck/pkg/plugin"
	"go.uber.org/zap"
)

func LogResultWithLogger(result *plugin.Result) {
	logger := zap.L()

	logger.Info(fmt.Sprintf("status: %s", result.Status))
	logger.Info(fmt.Sprintf("message: %s", result.StatusMessage))

	if result.Messages != nil {
		for _, msg := range result.Messages {
			logger.Info(msg)
		}
	}

	if result.Warnings != nil {
		for _, msg := range result.Warnings {
			logger.Warn(msg)
		}
	}

	if result.Errors != nil {
		for _, msg := range result.Errors {
			logger.Error(msg)
		}
	}
}