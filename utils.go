package god

import (
	"fmt"
	"sort"
)

func GetMessage(defaultMsg string, msgs ...string) string {
	if msgs != nil {
		return msgs[0]
	}
	return defaultMsg
}

func CommonPlaygroundValidation(tag string, messages []string, customFunc ...func(v interface{}) Schema) Validation {
	defaultMsg := fmt.Sprintf("Failed on %s validation!", tag)
	return Validation{
		Tag:     tag,
		Message: GetMessage(defaultMsg, messages...),
		Func: func(v interface{}) Schema {
			if customFunc != nil {
				return customFunc[0](v)
			}

			err := validate.Var(v, tag)
			return Schema{Error: err}
		},
	}
}

func CommonTypeValidation(v []Validation, value interface{}, typeName string, typeAssertFunc typeAssertFunc) Schema {
	schema := Schema{}
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
		schema.Error = nil
	}

	// Primeiro, verifica se o valor é do tipo correto
	_, ok = typeAssertFunc(value)
	if !ok {
		schema.Error = fmt.Errorf("value is not a %s!", typeName)
	}

	// Aplica as validações
	for _, validation := range v {
		// Aplica a validação
		err := validation.Func(value)
		if err.Error != nil {
			schema.Error = fmt.Errorf(validation.Message)
		}
	}

	return schema
}
