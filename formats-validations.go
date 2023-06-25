package god

import "fmt"

func Email(message ...string) Validation {
	return Validation{
		Tag: "email",
		Func: func(v interface{}) error {
			return validate.Var(v, "email")
		},
		Message: getMessage(message, "Failed on email validation!"),
	}
}

func Mongodb(message ...string) Validation {
	return Validation{
		Tag: "mongodb",
		Func: func(v interface{}) error {
			return validate.Var(v, "mongodb")
		},
		Message: getMessage(message, "Failed on mongodb validation!"),
	}
}

func Datetime(datetime string, message ...string) Validation {
	return Validation{
		Tag: "datetime",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("datetime=%s", datetime))
		},
		Message: getMessage(message, "Failed on datatime validation!"),
	}
}
