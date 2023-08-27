package god

import (
	"github.com/go-playground/validator/v10"
)

type Validation struct {
	Tag     string
	Func    SchemaFunc
	Message string
	Weight  int
}

var (
	validate = validator.New()
)

func Validate(v interface{}, f SchemaFunc) error {
	return f(v).Error
}
