package god

type SchemaFunc func(v interface{}) Schema

type Schema struct {
	Error  error
	Errors map[string]error
	Type   string
}
