package god

import (
	"fmt"
	"reflect"
)

type Map map[string]SchemaFunc

var structFieldCache = make(map[reflect.Type][]reflect.StructField)

func Object(m Map) SchemaFunc {
	return func(value interface{}) Schema {
		god_err := GodError{
			errors: make(map[string]interface{}),
		}

		rv := reflect.ValueOf(map[string]interface{}{})
		if value != nil {
			rv = reflect.ValueOf(value)
		}

		schema := Schema{Type: "Object"}

		switch rv.Kind() {
		case reflect.Map:
			objectMap(value, m, god_err)

		case reflect.Struct:
			objectStruct(value, m, god_err)

		case reflect.Ptr:
			value_of_pointer := structToMap(value)
			objectMap(value_of_pointer, m, god_err)

		default:
			err := fmt.Errorf("value is neither a map nor a struct. Is: %v", rv.Kind())
			schema.Error = err
		}

		if len(god_err.errors) > 0 {
			schema.Error = god_err
		}

		return schema
	}
}

func objectMap(value interface{}, m Map, god_err GodError) error {
	// Trate como um map se for um map
	mapValue, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("value is not a map")
	}

	if len(mapValue) == 0 {
		return nil
	}

	for key, validation := range m {
		fieldValue, _ := mapValue[key]
		if err := validation(fieldValue); err.Error != nil {
			god_err.errors[key] = err.Error
		}
	}

	return nil
}

func objectStruct(value interface{}, items Map, god_err GodError) error {
	v := reflect.ValueOf(value)
	t := v.Type()

	// Tente obter os campos do cache
	fields, ok := structFieldCache[t]
	if !ok {
		fields = make([]reflect.StructField, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			fields[i] = t.Field(i)
		}
		structFieldCache[t] = fields
	}

	if len(fields) == 0 {
		return nil
	}

	// Trate como uma struct se for uma struct
	for key, validation := range items {
		var fieldName string

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			queryTag := field.Tag.Get("query")
			godTag := field.Tag.Get("god")

			if jsonTag == key || queryTag == key || godTag == key {
				fieldName = field.Name
				break
			}
		}

		fieldValue := v.FieldByName(fieldName)
		if !fieldValue.IsValid() {
			god_err.errors[key] = fmt.Errorf("field not found in struct")
			continue
		}

		err := validation(fieldValue.Interface())
		if err.Error != nil {
			god_err.errors[key] = err.Error
		}
	}

	return nil
}

func structToMap(obj interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("json")
		if key == "" {
			key = field.Name
		}
		out[key] = v.Field(i).Interface()
	}
	return out
}
