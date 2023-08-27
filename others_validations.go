package god

import (
	"fmt"
	"reflect"
)

func Required(message ...string) Validation {
	return Validation{
		Weight:  100,
		Tag:     "required",
		Message: GetMessage("Failed on required validation!", message...),
		Func: func(v interface{}) Schema {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Slice && rv.Len() == 0 {
				err := fmt.Errorf(GetMessage("slice must have at least one element", message...))
				return Schema{Error: err}
			}

			err := validate.Var(v, "required")
			return Schema{Error: err}
		},
	}
}
