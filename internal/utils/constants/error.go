package constants

import "errors"

var (
	// global
	ErrUnexpected = errors.New("unexpected error")

	// not found
	ErrTaskNotFound = errors.New("task not found")

	// config
	ErrLoadConfig = errors.New("failed to load config file")
)
