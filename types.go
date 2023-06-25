package god

import (
	"fmt"
	"sort"
)

type typeAssertFunc func(interface{}) (interface{}, bool)

func commonValidation(v []Validation, value interface{}, typeName string, typeAssertFunc typeAssertFunc) error {
	// Ordena as validações por peso
	sort.Slice(v, func(i, j int) bool {
		return v[i].Weight > v[j].Weight
	})

	var isRequired bool
	if len(v) > 0 {
		isRequired = v[0].Tag == "required"
	}

	// Valida se é um campo opcional
	// Caso tenha nenhum valor deve ignorar as validações
	_, ok := typeAssertFunc(value)
	if !isRequired && !ok {
		return nil
	}

	// Primeiro, verifica se o valor é do tipo correto
	_, ok = typeAssertFunc(value)
	if !ok {
		return fmt.Errorf("value is not a %s!", typeName)
	}

	// Aplica as validações
	for _, validation := range v {
		// Aplica a validação
		err := validation.Func(value)
		if err != nil {
			return fmt.Errorf(validation.Message)
		}
	}

	return nil
}

// String for string type
func String(v ...Validation) Schema {
	return func(value interface{}) error {
		return commonValidation(v, value, "string", func(val interface{}) (interface{}, bool) {
			if out, ok := val.(string); ok {
				return out, ok
			}
			if out, ok := val.(*string); ok {
				return out, ok
			}
			return nil, false
		})
	}
}

// Float32 for float32 type
func Float32(v ...Validation) Schema {
	return func(value interface{}) error {
		return commonValidation(v, value, "float", func(val interface{}) (interface{}, bool) {
			if out, ok := val.(float32); ok {
				return out, ok
			}
			if out, ok := val.(*float32); ok {
				return out, ok
			}
			return nil, false
		})
	}
}

// Float64 for float64 type
func Float64(v ...Validation) Schema {
	return func(value interface{}) error {
		return commonValidation(v, value, "float64", func(val interface{}) (interface{}, bool) {
			if out, ok := val.(float64); ok {
				return out, ok
			}
			if out, ok := val.(*float64); ok {
				return out, ok
			}
			return nil, false
		})
	}
}

// Int for int32 type
func Int(v ...Validation) Schema {
	return func(value interface{}) error {
		return commonValidation(v, value, "int", func(val interface{}) (interface{}, bool) {
			if out, ok := val.(int); ok {
				return out, ok
			}
			if out, ok := val.(*int); ok {
				return out, ok
			}
			return nil, false
		})
	}
}

// Int64 for int64 type
func Int64(v ...Validation) Schema {
	return func(value interface{}) error {
		return commonValidation(v, value, "int64", func(val interface{}) (interface{}, bool) {
			if out, ok := val.(int64); ok {
				return out, ok
			}
			if out, ok := val.(*int64); ok {
				return out, ok
			}
			return nil, false
		})
	}
}

// Bool for bool type
func Bool(v ...Validation) Schema {
	return func(value interface{}) error {
		return commonValidation(v, value, "bool", func(val interface{}) (interface{}, bool) {
			if out, ok := val.(bool); ok {
				return out, ok
			}
			if out, ok := val.(*bool); ok {
				return out, ok
			}
			return nil, false
		})
	}
}
