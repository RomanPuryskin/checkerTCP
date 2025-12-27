package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(str interface{}) error {
	validator := validator.New()
	err := validator.Struct(str)
	if err != nil {
		return fmt.Errorf("[ValidateStruct]: %w", err)
	}
	return nil
}

func ValidatePort(port int) error {
	if port < 0 || port > 65535 {
		return errors.New("invalid port")
	}
	return nil
}

func ValidateTimeout(t string) error {
	_, err := time.ParseDuration(t)
	if err != nil {
		return err
	}

	return nil
}
