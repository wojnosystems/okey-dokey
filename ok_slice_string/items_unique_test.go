package ok_slice_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestItemsUnique_Validate(t *testing.T) {
	cases := map[string]struct {
		input    []optional.String
		expected string
	}{
		"ok": {
			input: []optional.String{
				optional.StringFrom("1"),
				optional.StringFrom("2"),
				optional.StringFrom("3"),
			},
			expected: "",
		},
		"duplicate": {
			input: []optional.String{
				optional.StringFrom("1"),
				optional.StringFrom("2"),
				optional.StringFrom("2"),
			},
			expected: "must have only unique elements, but had duplicates of 2",
		},
		"duplicates": {
			input: []optional.String{
				optional.StringFrom("1"),
				optional.StringFrom("2"),
				optional.StringFrom("2"),
				optional.StringFrom("3"),
				optional.StringFrom("4"),
				optional.StringFrom("3"),
			},
			expected: "must have only unique elements, but had duplicates of 2, 3",
		},
		"nil": {
			expected: "",
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := ItemsUnique{}
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
