package god

import (
	"fmt"
	"strings"
)

func Boolean(message ...string) Validation {
	return Validation{
		Tag: "boolean",
		Func: func(v interface{}) error {
			return validate.Var(v, "boolean")
		},
		Message: GetMessage(message, "Failed on boolean validation!"),
	}
}

func Lowercase(message ...string) Validation {
	return Validation{
		Tag: "lowercase",
		Func: func(v interface{}) error {
			return validate.Var(v, "lowercase")
		},
		Message: GetMessage(message, "Failed on lowercase validation!"),
	}
}

func Uppercase(message ...string) Validation {
	return Validation{
		Tag: "uppercase",
		Func: func(v interface{}) error {
			return validate.Var(v, "uppercase")
		},
		Message: GetMessage(message, "Failed on uppercase validation!"),
	}
}

func Contains(value string, message ...string) Validation {
	return Validation{
		Tag: "contains",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("contains=%s", value))
		},
		Message: GetMessage(message, fmt.Sprintf("field contains no value: %s", value)),
	}
}

func Number(message ...string) Validation {
	return Validation{
		Tag: "number",
		Func: func(v interface{}) error {
			return validate.Var(v, "number")
		},
		Message: GetMessage(message, "Failed on number validation!"),
	}
}

func OneOf(options []string, message ...string) Validation {
	return Validation{
		Tag: "oneof",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("oneof=%s", strings.Join(options, " ")))
		},
		Message: GetMessage(message, "Failed on oneOf validation!"),
	}
}
