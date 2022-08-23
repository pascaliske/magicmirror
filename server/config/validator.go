package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/viper"
)

/**
 * Validate config at given path.
 */
func validateConfig() (bool, error) {
	// variable to be validated
	var config Config

	// setup validator
	var validate *validator.Validate = validator.New()

	// prepare config for validation
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return false, errors.New("Unable to validate config.")
	}

	// validate config
	if err := validate.Struct(config); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				logger.Error("Required value for \"%s\" is missing.", err.Namespace())
			case "unique":
				logger.Error("Duplicate value found in \"%s\".", err.Namespace())
			case "oneof":
				logger.Error("Value for \"%s\" should be one of \"%s\" (actual: %s).", err.Namespace(), err.Param(), err.Value())
			case "url":
				logger.Error("Value for \"%s\" should be of type URL (actual: %s).", err.Namespace(), err.Value())
			case "startswith":
				logger.Error("Value for \"%s\" should start with \"%s\" (actual: %s).", err.Namespace(), err.Param(), err.Value())
			case "endsnotwith":
				logger.Error("Value for \"%s\" should not end with \"%s\" (actual: %s).", err.Namespace(), err.Param(), err.Value())
			}
		}

		// invalid config
		return false, errors.New("Provided config is invalid - please check validation errors above.")
	}

	// valid config
	return true, nil
}
