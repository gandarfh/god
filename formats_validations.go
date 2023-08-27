package god

import "fmt"

func Email(message ...string) Validation {
	return CommonPlaygroundValidation(
		"email",
		message,
	)
}

func Mongodb(message ...string) Validation {
	return CommonPlaygroundValidation(
		"mongodb",
		message,
	)
}

func Datetime(datetime string, message ...string) Validation {
	return CommonPlaygroundValidation(
		"datetime",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("datetime=%s", datetime))
			return Schema{Error: err}
		},
	)

}
