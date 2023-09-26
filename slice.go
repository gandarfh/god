package god

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Slice(vf SchemaFunc, v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		schema := Schema{}
		sort.Slice(v, func(i, j int) bool {
			return v[i].Weight > v[j].Weight
		})

		var isRequired bool
		if len(v) > 0 {
			isRequired = v[0].Tag == "required"
		}

		rv := reflect.ValueOf([]interface{}{})
		if value != nil {
			rv = reflect.ValueOf(value)
		}

		if !isRequired && rv.Len() == 0 {
			schema.Error = nil
		}

		if rv.Kind() != reflect.Slice {
			schema.Error = fmt.Errorf("value is not a slice")
		}

		var errors MultiError

		for _, validation := range v {
			err := validation.Func(value)
			if err.Error != nil {
				errors = append(errors, fmt.Errorf(validation.Message))
			}
		}

		for i := 0; i < rv.Len(); i++ {
			item := rv.Index(i).Interface()

			if err := vf(item); err.Error != nil {
				errors = append(errors, fmt.Errorf("index %d: %v", i, err))
			}
		}

		if len(errors) > 0 {
			schema.Error = errors
		}

		return schema
	}
}

type MultiError []error

func (m MultiError) Error() string {
	errs := make([]string, len(m))
	for i, err := range m {
		errs[i] = err.Error()
	}
	return strings.Join(errs, "; ")
}

func (m MultiError) Errors() []error {
	return m
}
