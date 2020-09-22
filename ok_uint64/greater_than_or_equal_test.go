package ok_uint64

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestGreaterThanOrEqual_Validate(t *testing.T) {
	cases := map[string]struct {
		greaterThanOrEq uint64
		input           optional.Uint64
		expected        string
	}{
		"ok": {
			greaterThanOrEq: 3,
			input:           optional.Uint64From(4),
			expected:        "",
		},
		"too large": {
			greaterThanOrEq: 6,
			input:           optional.Uint64From(5),
			expected:        "must be greater than or equal to 6",
		},
		"equal": {
			greaterThanOrEq: 5,
			input:           optional.Uint64From(5),
			expected:        "",
		},
		"nil": {
			greaterThanOrEq: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := GreaterThanOrEqual{
				Value: c.greaterThanOrEq,
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
