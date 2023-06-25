package god

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Slice(vf Schema, v ...Validation) Schema {
	return func(value interface{}) error {
		// Ordena as validações por peso
		sort.Slice(v, func(i, j int) bool {
			return v[i].Weight > v[j].Weight
		})

		rv := reflect.ValueOf(value)
		var isRequired bool
		if len(v) > 0 {
			isRequired = v[0].Tag == "required"
		}

		if !isRequired && rv.Len() == 0 {
			return nil
		}

		// Valida se é um slice
		if rv.Kind() != reflect.Slice {
			return fmt.Errorf("value is not a slice")
		}

		var errors MultiError

		// Aplica as validações
		for _, validation := range v {
			// Aplica a validação
			err := validation.Func(value)
			if err != nil {
				errors = append(errors, fmt.Errorf(validation.Message))
			}
		}

		// Aplica a função de validação a cada item na fatia.
		for i := 0; i < rv.Len(); i++ {
			item := rv.Index(i).Interface()

			if err := vf(item); err != nil {
				errors = append(errors, fmt.Errorf("index %d: %v", i, err))
			}
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
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
