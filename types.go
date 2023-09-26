package god

type typeAssertFunc func(interface{}) (interface{}, bool)

// String for string type
func String(v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		return CommonTypeValidation(v, value, "string", func(val interface{}) (interface{}, bool) {
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
func Float32(v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		return CommonTypeValidation(v, value, "float", func(val interface{}) (interface{}, bool) {
			if out, ok := ConvertToNumeric(value); ok {
				return out.ToFloat32(), ok
			}

			return nil, false
		})
	}
}

// Float64 for float64 type
func Float64(v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		return CommonTypeValidation(v, value, "float64", func(val interface{}) (interface{}, bool) {
			if out, ok := ConvertToNumeric(value); ok {
				return out.ToFloat64(), ok
			}

			return nil, false
		})
	}
}

// Int for int32 type
func Int(v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		return CommonTypeValidation(v, value, "int", func(val interface{}) (interface{}, bool) {
			if out, ok := ConvertToNumeric(value); ok {
				return out.ToInt32(), ok
			}

			return nil, false
		})
	}
}

// Int64 for int64 type
func Int64(v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		return CommonTypeValidation(v, value, "int64", func(val interface{}) (interface{}, bool) {
			if out, ok := ConvertToNumeric(value); ok {
				return out.ToInt64(), ok
			}

			return nil, false
		})
	}
}

// Bool for bool type
func Bool(v ...Validation) SchemaFunc {
	return func(value interface{}) Schema {
		return CommonTypeValidation(v, value, "bool", func(val interface{}) (interface{}, bool) {
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
