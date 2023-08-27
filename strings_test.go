package god_test

import (
	"testing"

	"github.com/gandarfh/god"
)

func TestStringValidations(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		validation god.Validation
		value      string
		expectErr  bool
	}{
		{
			name:       "Boolean",
			validation: god.Boolean("Failed on boolean validation!"),
			value:      "true",
			expectErr:  false,
		},
		{
			name:       "Lowercase",
			validation: god.Lowercase("Failed on lowercase validation!"),
			value:      "lowercase",
			expectErr:  false,
		},
		{
			name:       "Uppercase",
			validation: god.Uppercase("Failed on uppercase validation!"),
			value:      "UPPERCASE",
			expectErr:  false,
		},
		{
			name:       "Contains",
			validation: god.Contains("test", "Field contains no value: test"),
			value:      "this is a test",
			expectErr:  false,
		},
		{
			name:       "Number",
			validation: god.Number("Failed on number validation!"),
			value:      "1234",
			expectErr:  false,
		},
		{
			name:       "BooleanError",
			validation: god.Boolean("Failed on boolean validation!"),
			value:      "notboolean",
			expectErr:  true,
		},
		{
			name:       "LowercaseError",
			validation: god.Lowercase("Failed on lowercase validation!"),
			value:      "NOTLOWERCASE",
			expectErr:  true,
		},
		{
			name:       "UppercaseError",
			validation: god.Uppercase("Failed on uppercase validation!"),
			value:      "notuppercase",
			expectErr:  true,
		},
		{
			name:       "ContainsError",
			validation: god.Contains("test", "Field contains no value: test"),
			value:      "does not contain",
			expectErr:  true,
		},
		{
			name:       "NumberError",
			validation: god.Number("Failed on number validation!"),
			value:      "notnumber",
			expectErr:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.validation.Func(testCase.value)
			if testCase.expectErr && err.Error == nil {
				t.Errorf("Expected error, got nil")
			}
			if !testCase.expectErr && err.Error != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
