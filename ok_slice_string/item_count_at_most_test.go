package ok_slice_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestItemCountAtMost_Validate(t *testing.T) {
	cases := map[string]struct {
		atMost   int
		input    []optional.String
		expected string
	}{
		"ok": {
			atMost: 4,
			input: []optional.String{
				optional.StringFrom("1"),
				optional.StringFrom("2"),
				optional.StringFrom("3"),
			},
			expected: "",
		},
		"too many": {
			atMost: 2,
			input: []optional.String{
				optional.StringFrom("1"),
				optional.StringFrom("2"),
				optional.StringFrom("3"),
			},
			expected: "cannot have more than 2 items",
		},
		"nil": {
			atMost: 6,
		},
		"nil but expecting 0 items": {
			atMost: 0,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := ItemCountAtMost{
				AtMost: c.atMost,
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