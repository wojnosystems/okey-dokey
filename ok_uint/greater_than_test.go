package ok_uint

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestGreaterThan_Validate(t *testing.T) {
	cases := map[string]struct {
		greaterThan uint
		input       optional.Uint
		expected    string
	}{
		"ok": {
			greaterThan: 3,
			input:       optional.UintFrom(4),
			expected:    "",
		},
		"too large": {
			greaterThan: 6,
			input:       optional.UintFrom(5),
			expected:    "must be greater than 6",
		},
		"equal": {
			greaterThan: 5,
			input:       optional.UintFrom(5),
			expected:    "must be greater than 5",
		},
		"nil": {
			greaterThan: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := GreaterThan{
				Value: c.greaterThan,
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
