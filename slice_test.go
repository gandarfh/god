package god_test

import (
	"fmt"
	"testing"

	"github.com/gandarfh/god"
)

func TestSlice(t *testing.T) {
	t.Run("it should validate each item in the slice", func(t *testing.T) {
		vf := func(value interface{}) god.Schema {
			schema := god.Schema{}
			v, ok := value.(string)
			if !ok {
				schema.Error = fmt.Errorf("value is not a string")
			}
			if v != "ok" {
				schema.Error = fmt.Errorf("value is not 'ok'")
			}
			return schema
		}

		slice := []string{"ok", "not ok", "ok"}
		err := god.Slice(vf)(slice)
		if err.Error == nil {
			t.Fatal("expected error but got none")
		}

		me, ok := err.Error.(god.MultiError)
		if !ok {
			t.Fatalf("expected MultiError but got %T", err)
		}
		if len(me) != 1 {
			t.Fatalf("expected 1 error but got %d", len(me))
		}
		if me[0].Error() != "index 1: {value is not 'ok' map[] }" {
			t.Fatalf("unexpected error message: %v", me[0])
		}
	})

	t.Run("it should validate strings into slice", func(t *testing.T) {

		slice := []interface{}{"test1", "test2", "test3"}

		// Usando a função Slice com a função de validação.
		sliceSchema := god.Slice(
			god.String(god.Required("String is required")),
			god.Required("Slice is required"),
		)

		// Validando o slice.
		if err := god.Validate(slice, sliceSchema); err != nil {
			t.Errorf("Validation failed: %v", err)
		}
	})

	t.Run("it should validate ints into slice", func(t *testing.T) {
		slice := []interface{}{5, 10, 15}

		// Usando a função Slice com a função de validação.
		sliceSchema := god.Slice(
			god.Int(god.Required("Int is required")),
			god.Required("Slice is required"),
		)

		// Validando o slice.
		if err := god.Validate(slice, sliceSchema); err != nil {
			t.Errorf("Validation failed: %v", err)
		}
	})
}
