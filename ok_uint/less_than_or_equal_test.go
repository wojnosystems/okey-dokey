package ok_uint

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"okey-dokey/bad"
	"testing"
)

func TestLessThanOrEqual_Validate(t *testing.T) {
	cases := map[string]struct {
		lessThanOrEq uint
		input        *uint
		expected     string
	}{
		"ok": {
			lessThanOrEq: 3,
			input:        addrOf(2),
			expected:     "",
		},
		"too large": {
			lessThanOrEq: 6,
			input:        addrOf(7),
			expected:     "must be less than or equal to 6",
		},
		"equal": {
			lessThanOrEq: 5,
			input:        addrOf(5),
			expected:     "",
		},
		"nil": {
			lessThanOrEq: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LessThanOrEqual{
				Value: c.lessThanOrEq,
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
