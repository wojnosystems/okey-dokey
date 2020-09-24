package ok_slice_uint64

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_range"
	"testing"
)

func TestItemCountBetween_Validate(t *testing.T) {
	cases := map[string]struct {
		atLeast  int
		atMost   int
		input    []optional.Uint64
		expected string
	}{
		"ok": {
			atLeast: 2,
			atMost:  4,
			input: []optional.Uint64{
				optional.Uint64From("1"),
				optional.Uint64From("2"),
				optional.Uint64From("3"),
			},
			expected: "",
		},
		"exact": {
			atLeast: 3,
			atMost:  3,
			input: []optional.Uint64{
				optional.Uint64From("1"),
				optional.Uint64From("2"),
				optional.Uint64From("3"),
			},
			expected: "",
		},
		"too rew": {
			atLeast: 2,
			atMost:  4,
			input: []optional.Uint64{
				optional.Uint64From("1"),
			},
			expected: "must have between 2 and 4 items",
		},
		"too many": {
			atLeast: 2,
			atMost:  4,
			input: []optional.Uint64{
				optional.Uint64From("1"),
				optional.Uint64From("2"),
				optional.Uint64From("3"),
				optional.Uint64From("4"),
				optional.Uint64From("5"),
			},
			expected: "must have between 2 and 4 items",
		},
		"nil": {
			atLeast:  2,
			atMost:   4,
			expected: "must have between 2 and 4 items",
		},
		"nil but expecting 0 items": {
			atLeast: 0,
			atMost:  0,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := ItemCountBetween{
				Between: ok_range.IntBetween(c.atLeast, c.atMost),
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