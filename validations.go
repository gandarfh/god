package god

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Schema func(v interface{}) error

type Validation struct {
	Tag     string
	Func    Schema
	Message string
	Weight  int
}

var (
	validate = validator.New()
)

func Validate(v interface{}, f Schema) error {
	return f(v)
}

func Required(message ...string) Validation {
	return Validation{
		Tag: "required",
		Func: func(v interface{}) error {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Slice && rv.Len() == 0 {
				return fmt.Errorf(GetMessage(message, "slice must have at least one element"))
			}

			return validate.Var(v, "required")
		},
		Message: GetMessage(message, "Failed on required validation!"),
		Weight:  100,
	}
}

func Min(size int, message ...string) Validation {
	return Validation{
		Tag: "min",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("min=%d", size))
		},
		Message: GetMessage(message, fmt.Sprintf("the value is less than %d", size)),
	}
}

func Max(size int, message ...string) Validation {
	return Validation{
		Tag: "max",
		Func: func(v interface{}) error {
			return validate.Var(v, fmt.Sprintf("max=%d", size))
		},
		Message: GetMessage(message, fmt.Sprintf("the value is greater than %d", size)),
	}
}

func GetMessage(msg []string, defaultMsg string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return defaultMsg
}
