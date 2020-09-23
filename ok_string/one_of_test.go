package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestOneOf_Validate(t *testing.T) {
	cases := map[string]struct {
		oneOf    SortedSet
		input    optional.String
		expected string
	}{
		"ok": {
			oneOf:    (&SortedSetBuilder{}).Add("apple").Add("banana").Add("carrot").Add("durian").Sort(),
			input:    optional.StringFrom("carrot"),
			expected: "",
		},
		"missing": {
			oneOf:    (&SortedSetBuilder{}).Add("apple").Add("banana").Add("carrot").Add("durian").Sort(),
			input:    optional.StringFrom("invalid"),
			expected: "must be one of the following: apple, banana, carrot, durian",
		},
		"missing long list": {
			oneOf: (&SortedSetBuilder{}).
				Add("apple").
				Add("banana").
				Add("carrot").
				Add("durian").
				Add("eggplant").
				Add("fig").
				Add("grape").
				Add("honeydew").
				Add("ifood").
				Add("jackfruit").
				Add("kiwi").
				Sort(),
			input:    optional.StringFrom("invalid"),
			expected: "must be one of the following: apple, banana, carrot, durian, eggplant, fig, grape, honeydew, ifood, jackfruit, ...",
		},
		"empty": {
			oneOf: (&SortedSetBuilder{}).Sort(),
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
