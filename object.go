package god

import (
	"fmt"
	"reflect"
)

type Map map[string]Schema

func Object(m Map) Schema {
	return func(value interface{}) error {
		errors := make(map[string]error)

		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.Map:
			objectMap(value, m, errors)

		case reflect.Struct:
			objectStruct(value, m, errors)

		default:
			return fmt.Errorf("value is neither a map nor a struct")
		}

		if len(errors) > 0 {
			return fmt.Errorf("%v", errors)
		}

		return nil
	}
}

func objectMap(value interface{}, m Map, errors map[string]error) error {
	// Trate como um map se for um map
	mapValue, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("value is not a map")
	}

	for field, validation := range m {
		fieldValue, _ := mapValue[field]
		if err := validation(fieldValue); err != nil {
			errors[field] = err
		}
	}

	return nil
}

func objectStruct(value interface{}, m Map, errors map[string]error) error {
	v := reflect.ValueOf(value)
	t := v.Type()

	// Trate como uma struct se for uma struct
	for key, validation := range m {
		var fieldName string
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == key {
				fieldName = field.Name
				break
			}
		}

		if fieldName == "" {
			errors[key] = fmt.Errorf("unknown field")
			continue
		}

		fieldValue := v.FieldByName(fieldName)
		if !fieldValue.IsValid() {
			errors[key] = fmt.Errorf("field not found in struct")
			continue
		}

		err := validation(fieldValue.Interface())
		if err != nil {
			errors[key] = err
		}
	}

	return nil
}
