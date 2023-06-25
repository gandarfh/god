package god

import "fmt"

func Eq(value interface{}, message ...string) Validation {
	return Validation{
		Tag: "eq",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("eq=%v", value))
		},
		Message: getMessage(message, fmt.Sprintf("the value is not equal %v", value)),
	}
}

func Ne(value interface{}, message ...string) Validation {
	return Validation{
		Tag: "ne",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("ne=%v", value))
		},
		Message: getMessage(message, fmt.Sprintf("the value is equal %v", value)),
	}
}

func Gt(value int, message ...string) Validation {
	return Validation{
		Tag: "gt",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("gt=%d", value))
		},
		Message: getMessage(message, fmt.Sprintf("the value is not greater than %d", value)),
	}
}

func Gte(value int, message ...string) Validation {
	return Validation{
		Tag: "gte",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("gte=%d", value))
		},
		Message: getMessage(message, fmt.Sprintf("the value is not greater than or equal %d", value)),
	}
}

func Lt(value int, message ...string) Validation {
	return Validation{
		Tag: "lt",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("lt=%d", value))
		},
		Message: getMessage(message, fmt.Sprintf("the value is not less than %d", value)),
	}
}

func Lte(value int, message ...string) Validation {
	return Validation{
		Tag: "lte",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("lte=%d", value))
		},
		Message: getMessage(message, fmt.Sprintf("the value is not less than or equal %d", value)),
	}
}
