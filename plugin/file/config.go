package filePlugin

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Path   string
	Exists bool
}

func validateConfig(config *Config) error {
	return validation.ValidateStruct(config,
		validation.Field(&config.Path, validation.Required),
	)
}
