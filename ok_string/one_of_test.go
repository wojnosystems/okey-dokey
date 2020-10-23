package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	sorted_set "github.com/wojnosystems/go-sorted-set"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestOneOf_Validate(t *testing.T) {
	cases := map[string]struct {
		oneOf    sorted_set.String
		input    optional.String
		expected string
	}{
		"ok": {
			oneOf:    sorted_set.NewString("apple", "durian", "carrot", "banana").Sort(),
			input:    optional.StringFrom("carrot"),
			expected: "",
		},
		"missing": {
			oneOf:    sorted_set.NewString("apple", "durian", "carrot", "banana").Sort(),
			input:    optional.StringFrom("invalid"),
			expected: "must be one of the following: apple, banana, carrot, durian",
		},
		"missing long list": {
			oneOf: sorted_set.NewString(
				"apple",
				"banana",
				"carrot",
				"durian",
				"eggplant",
				"fig",
				"grape",
				"honeydew",
				"ifood",
				"jackfruit",
				"kiwi").Sort(),
			input:    optional.StringFrom("invalid"),
			expected: "must be one of the following: apple, banana, carrot, durian, eggplant, fig, grape, honeydew, ifood, jackfruit, ...",
		},
		"empty": {
			oneOf: sorted_set.NewString().Sort(),
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := OneOf{
				Only: c.oneOf,
			}
			validationErrors := bad.SliceMessageReceiver{}
			validator.Validate(c.input, &validationErrors)
			if len(c.expected) != 0 {
				require.Len(t, validationErrors, 1)
				assert.Equal(t, c.expected, validationErrors[0])
			} else {
				assert.Equal(t, 0, len(validationErrors))
			}
		})
	}
}
