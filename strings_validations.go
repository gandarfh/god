package god

import (
	"fmt"
	"strings"
)

func Boolean(message ...string) Validation {
	return CommonPlaygroundValidation(
		"boolean",
		message,
	)
}

func Lowercase(message ...string) Validation {
	return CommonPlaygroundValidation(
		"lowercase",
		message,
	)
}

func Uppercase(message ...string) Validation {
	return CommonPlaygroundValidation(
		"uppercase",
		message,
	)
}

func Contains(value string, message ...string) Validation {
	return CommonPlaygroundValidation(
		"uppercase",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("contains=%s", value))
			return Schema{Error: err}
		},
	)
}

func Number(message ...string) Validation {
	return CommonPlaygroundValidation(
		"number",
		message,
	)
}

func OneOf(options []string, message ...string) Validation {
	return CommonPlaygroundValidation(
		"oneof",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("oneof=%s", strings.Join(options, " ")))
			return Schema{Error: err}
		},
	)
}
