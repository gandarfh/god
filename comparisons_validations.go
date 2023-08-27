package god

import "fmt"

func Eq(value interface{}, message ...string) Validation {
	return CommonPlaygroundValidation(
		"eq",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("eq=%v", value))
			return Schema{Error: err}
		},
	)

}

func Ne(value interface{}, message ...string) Validation {
	return CommonPlaygroundValidation(
		"ne",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("ne=%v", value))
			return Schema{Error: err}
		},
	)
}

func Gt(value int, message ...string) Validation {
	return CommonPlaygroundValidation(
		"gt",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("gt=%d", value))
			return Schema{Error: err}
		},
	)
}

func Gte(value int, message ...string) Validation {
	return CommonPlaygroundValidation(
		"gte",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("gte=%d", value))
			return Schema{Error: err}
		},
	)
}

func Lt(value int, message ...string) Validation {
	return CommonPlaygroundValidation(
		"lt",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("lt=%d", value))
			return Schema{Error: err}
		},
	)
}

func Lte(value int, message ...string) Validation {
	return CommonPlaygroundValidation(
		"lte",
		message,
		func(v interface{}) Schema {
			err := validate.Var(v, fmt.Sprintf("lte=%d", value))
			return Schema{Error: err}
		},
	)
}
