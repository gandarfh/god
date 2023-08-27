package god

import (
	"bytes"
	"fmt"
	"strings"
)

type GodError struct {
	errors map[string]interface{}
}

func (err GodError) Error() string {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "{")
	var parts []string
	for key, value := range err.errors {
		switch v := value.(type) {
		case string:
			parts = append(parts, fmt.Sprintf("\"%s\": \"%s\"", key, v))
		case error:
			if nestedErr, ok := v.(GodError); ok {
				// Se for um GodError aninhado, serializamos ele recursivamente
				parts = append(parts, fmt.Sprintf("\"%s\": %s", key, nestedErr.Error()))
			} else {
				parts = append(parts, fmt.Sprintf("\"%s\": \"%s\"", key, v.Error()))
			}
		}
	}
	fmt.Fprintf(b, strings.Join(parts, ","))
	fmt.Fprintf(b, "}")
	return b.String()
}

func ErrorsToMap(errors error) map[string]interface{} {
	fields := map[string]interface{}{}

	for key, err := range errors.(GodError).errors {
		switch v := err.(type) {
		case string:
			fields[key] = v
		case error:
			if godErr, ok := v.(GodError); ok {
				fields[key] = ErrorsToMap(godErr)
			} else {
				fields[key] = v.Error()
			}
		}
	}

	return fields
}
