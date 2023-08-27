package god

import "fmt"

func Min(size int, message ...string) Validation {
	message = append([]string{fmt.Sprintf("the value is less than %d", size)}, message...)
	return CommonPlaygroundValidation(
		"min",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("min=%d", size))
			return Schema{Error: err}
		},
	)
}

func Max(size int, message ...string) Validation {
	message = append([]string{fmt.Sprintf("the value is greater than %d", size)}, message...)
	return CommonPlaygroundValidation(
		"max",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("max=%d", size))
			return Schema{Error: err}
		},
	)
}
