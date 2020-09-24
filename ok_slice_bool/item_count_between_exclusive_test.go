package ok_slice_bool

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_range"
	"testing"
)

func TestItemCountBetweenExclusive_Validate(t *testing.T) {
	cases := map[string]struct {
		atLeast  int
		atMost   int
		input    []optional.Bool
		expected string
	}{
		"ok": {
			atLeast: 2,
			atMost:  4,
			input: []optional.Bool{
				optional.BoolFrom("1"),
				optional.BoolFrom("2"),
				optional.BoolFrom("3"),
			},
			expected: "",
		},
		"too many": {
			atLeast: 2,
			atMost:  4,
			input: []optional.Bool{
				optional.BoolFrom("1"),
				optional.BoolFrom("2"),
				optional.BoolFrom("3"),
				optional.BoolFrom("4"),
				optional.BoolFrom("5"),
			},
			expected: "must have more than 2 but fewer than 4 items",
		},
		"too few": {
			atLeast: 2,
			atMost:  4,
			input: []optional.Bool{
				optional.BoolFrom("1"),
				optional.BoolFrom("2"),
			},
			expected: "must have more than 2 but fewer than 4 items",
		},
		"nil": {
			atLeast:  2,
			atMost:   4,
			expected: "must have more than 2 but fewer than 4 items",
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := ItemCountBetweenExclusive{
				Between: ok_range.MustIntBetweenExclusive(c.atLeast, c.atMost),
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
