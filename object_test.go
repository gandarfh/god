package god_test

import (
	"testing"

	"github.com/gandarfh/god"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestObject(t *testing.T) {
	userSchema := god.Object(god.Map{
		"name":  god.String(god.Required("Name is required")),
		"email": god.String(god.Required("Email is required"), god.Email("Email format is incorrect")),
	})

	t.Run("Struct: test with valid user", func(t *testing.T) {
		user := User{
			Name:  "John",
			Email: "john@example.com",
		}

		err := userSchema(user)
		if err.Error != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Struct: test with invalid user", func(t *testing.T) {
		user := User{
			Name:  "John",
			Email: "john", // Email inválido
		}

		err := userSchema(user)
		if err.Error == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Struct: test with non-struct value", func(t *testing.T) {
		value := "not a struct"

		err := userSchema(value)
		if err.Error == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Map: test with valid user", func(t *testing.T) {
		user := map[string]interface{}{
			"name":  "John",
			"email": "john@example.com",
		}

		err := userSchema(user)
		if err.Error != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Map: test with invalid user", func(t *testing.T) {
		user := map[string]interface{}{
			"name":  "John",
			"email": "john",
		}

		err := userSchema(user)
		if err.Error == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Map: test with non-map value", func(t *testing.T) {
		value := "not a map"

		err := userSchema(value)
		if err.Error == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

type TestFields struct {
	StringField string  `json:"stringField"`
	IntField    int64   `json:"intField"`
	BoolField   bool    `json:"boolField"`
	FloatField  float64 `json:"floatField"`
	SliceField  []int64 `json:"sliceField"`
	NestedField Nested  `json:"nestedField"`
}

type Nested struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestObjectValidation(t *testing.T) {
	t.Parallel()

	// Defining the validation schema
	schema := god.Object(god.Map{
		"stringField": god.String(god.Required(), god.Uppercase()),
		"intField":    god.Int64(god.Required()),
		"boolField":   god.Bool(),
		"floatField":  god.Float64(),
		"sliceField":  god.Slice(god.Int64(god.Required())),
		"nestedField": god.Object(god.Map{
			"name":  god.String(god.Required()),
			"email": god.String(god.Required(), god.Contains("@"), god.Email()),
		}),
	})

	// Create a test object
	testObj := TestFields{
		StringField: "TEST",
		IntField:    10,
		BoolField:   true,
		FloatField:  3.14,
		SliceField:  []int64{1, 2, 3},
		NestedField: Nested{
			Name:  "John",
			Email: "john@example.com",
		},
	}

	// Perform the validation
	if err := schema(testObj); err.Error != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Create a test map
	testMap := map[string]interface{}{
		"stringField": "TEST",
		"intField":    int64(10),
		"boolField":   true,
		"floatField":  3.14,
		"sliceField":  []int64{1, 2, 3},
		"nestedField": map[string]interface{}{
			"name":  "John",
			"email": "john@example.com",
		},
	}

	// Perform the validation
	if err := schema(testMap); err.Error != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
