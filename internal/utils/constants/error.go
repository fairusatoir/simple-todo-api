package constants

import "errors"

var (
	// global
	ErrUnexpected = errors.New("unexpected error")

	// not found
	ErrTaskNotFound = errors.New("task not found")

	// config
	ErrLoadConfig = errors.New("failed to load config file")

	// http global
	Err404 = errors.New("data not found")

	// http prod
	Err500Prod = errors.New("an internal server error has occurred. please try again later")
)
