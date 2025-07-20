package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pascaliske/magicmirror/logger"
)

/**
 * Validate config at given path.
 */
func (c *Config) validateConfig() (bool, error) {
	// setup validator
	var validate = validator.New()

	// prepare config for validation
	if err := c.viper.Unmarshal(&c); err != nil {
		fmt.Println(err)
		return false, errors.New("unable to validate config")
	}

	// validate config
	if err := validate.Struct(c); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required", "required_unless":
				logger.Error("Required value for \"%s\" is missing", err.Namespace())
			case "unique":
				logger.Error("Duplicate value found in \"%s\"", err.Namespace())
			case "oneof":
				logger.Error("Value for \"%s\" should be one of \"%s\" (actual: %s)", err.Namespace(), err.Param(), err.Value())
			case "url":
				logger.Error("Value for \"%s\" should be of type URL (actual: %s)", err.Namespace(), err.Value())
			case "number":
				logger.Error("Value for \"%s\" should be of type Number (actual: %s)", err.Namespace(), err.Value())
			case "hostname_rfc1123":
				logger.Error("Value for \"%s\" should be of type Hostname (actual: %s)", err.Namespace(), err.Value())
			case "startswith":
				logger.Error("Value for \"%s\" should start with \"%s\" (actual: %s)", err.Namespace(), err.Param(), err.Value())
			case "endsnotwith":
				logger.Error("Value for \"%s\" should not end with \"%s\" (actual: %s)", err.Namespace(), err.Param(), err.Value())
			case "min", "gte":
				logger.Error("Value for \"%s\" should be at least of length/size \"%s\" (actual: %s)", err.Namespace(), err.Param(), err.Value())
			default:
				logger.Error("Unknown error for \"%s\" (actual: %s)", err.Namespace(), err.Value())
			}
		}

		// invalid config
		return false, errors.New("provided config is invalid - please check validation errors above")
	}

	// valid config
	return true, nil
}
