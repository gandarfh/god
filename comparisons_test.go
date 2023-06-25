package god_test

import (
	"testing"

	"github.com/gandarfh/god"
)

func TestValidationFunctions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		validation god.Validation
		value      interface{}
		expectErr  bool
	}{
		{
			name:       "Test Eq Validation Success",
			validation: god.Eq(10),
			value:      10,
			expectErr:  false,
		},
		{
			name:       "Test Eq Validation Failure",
			validation: god.Eq(10),
			value:      20,
			expectErr:  true,
		},
		{
			name:       "Test Ne Validation Success",
			validation: god.Ne(10),
			value:      20,
			expectErr:  false,
		},
		{
			name:       "Test Ne Validation Failure",
			validation: god.Ne(10),
			value:      10,
			expectErr:  true,
		},
		{
			name:       "Test Gt Validation Success",
			validation: god.Gt(10),
			value:      20,
			expectErr:  false,
		},
		{
			name:       "Test Gt Validation Failure",
			validation: god.Gt(10),
			value:      5,
			expectErr:  true,
		},
		{
			name:       "Test Gte Validation Success",
			validation: god.Gte(10),
			value:      10,
			expectErr:  false,
		},
		{
			name:       "Test Gte Validation Failure",
			validation: god.Gte(10),
			value:      5,
			expectErr:  true,
		},
		{
			name:       "Test Lt Validation Success",
			validation: god.Lt(10),
			value:      5,
			expectErr:  false,
		},
		{
			name:       "Test Lt Validation Failure",
			validation: god.Lt(10),
			value:      20,
			expectErr:  true,
		},
		{
			name:       "Test Lte Validation Success",
			validation: god.Lte(10),
			value:      10,
			expectErr:  false,
		},
		{
			name:       "Test Lte Validation Failure",
			validation: god.Lte(10),
			value:      20,
			expectErr:  true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.validation.Func(testCase.value)
			if testCase.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !testCase.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
