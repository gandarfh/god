package god

import "reflect"

type Numeric interface {
	ToFloat64() float64
	ToFloat32() float32
	ToInt64() int64
	ToInt32() int
}

func ConvertToNumeric(value interface{}) (Numeric, bool) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		elemType := v.Elem().Type()

		switch elemType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return MyInt(v.Elem().Int()), true
		case reflect.Float32, reflect.Float64:
			return MyFloat64(v.Elem().Float()), true
		}
	} else {
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return MyInt(v.Int()), true
		case reflect.Float32, reflect.Float64:
			return MyFloat64(v.Float()), true
		}
	}

	return nil, false
}

type MyInt int

func (mi MyInt) ToFloat64() float64 {
	return float64(mi)
}

func (mf MyInt) ToFloat32() float32 {
	return float32(mf)
}

func (mi MyInt) ToInt32() int {
	return int(mi)
}

func (mi MyInt) ToInt64() int64 {
	return int64(mi)
}

type MyFloat64 float64

func (mf MyFloat64) ToFloat64() float64 {
	return float64(mf)
}

func (mf MyFloat64) ToFloat32() float32 {
	return float32(mf)
}

func (mf MyFloat64) ToInt32() int {
	return int(mf)
}

func (mi MyFloat64) ToInt64() int64 {
	return int64(mi)
}
