package ok_slice_int32

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestItemCountAtLeast_Validate(t *testing.T) {
	cases := map[string]struct {
		atLeast  int
		input    []optional.Int32
		expected string
	}{
		"ok": {
			atLeast: 2,
			input: []optional.Int32{
				optional.Int32From("1"),
				optional.Int32From("2"),
				optional.Int32From("3"),
			},
			expected: "",
		},
		"too short": {
			atLeast: 4,
			input: []optional.Int32{
				optional.Int32From("1"),
				optional.Int32From("2"),
				optional.Int32From("3"),
			},
			expected: "cannot have fewer than 4 items",
		},
		"nil": {
			atLeast:  6,
			expected: "cannot have fewer than 6 items",
		},
		"nil but expecting 0 items": {
			atLeast:  0,
			expected: "",
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := ItemCountAtLeast{
				AtLeast: c.atLeast,
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
