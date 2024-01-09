package pkg_validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

var mapHelepr = map[string]string{
	"required":  "is a required field",
	"email":     "is not a valid email address",
	"lowercase": "must contain at least one lowercase letter",
	"uppercase": "must contain at least one uppercase letter",
	"numeric":   "must contain at least one digit",
}

var needParam = []string{"min", "max", "containsany"}

func ValidatePayloads(err error) error {
	var field, param, value, tag, message string

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field = e.Field()
			tag = e.Tag()
			value = e.Value().(string)
			param = e.Param()

			if slices.Contains(needParam, tag) {
				message = errWithParam(field, value, tag, param)
				continue
			}

			if value != "" {
				value = fmt.Sprintf("'%s' ", value)
			}
			message = fmt.Sprintf("%s: %s%s", strings.ToLower(field), value, mapHelepr[tag])
		}

		return errors.New(message)
	}

	return nil
}

func errWithParam(field, value, tag, param string) string {
	var message string
	switch tag {
	case "min":
		message = fmt.Sprintf("must be at least %s characters long", param)
	case "max":
		message = fmt.Sprintf("must be less than %s characters", param)
	case "containsany":
		message = fmt.Sprintf("must contain at least one symbol of '%s'", param)
	}

	return fmt.Sprintf("%s: '%s' %s", field, value, message)
}
